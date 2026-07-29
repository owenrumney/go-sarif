package merge_test

import (
	"testing"

	"github.com/owenrumney/go-sarif/v3/pkg/report/merge"
	v210 "github.com/owenrumney/go-sarif/v3/pkg/report/v210/sarif"
	v22 "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newV22Report(t *testing.T, toolName string, ruleIDs ...string) *v22.Report {
	t.Helper()

	report := v22.NewReport()
	run := v22.NewRunWithInformationURI(toolName, "https://example.com/"+toolName)

	for _, ruleID := range ruleIDs {
		run.AddRule(ruleID).WithDescription("description of " + ruleID)
		run.AddDistinctArtifact("file:///code/" + ruleID + ".tf")
		run.CreateResultForRule(ruleID).
			WithLevel("error").
			WithMessage(v22.NewTextMessage(ruleID + " was breached"))
	}

	report.AddRun(run)
	return report
}

func newV210Report(t *testing.T, toolName string, ruleIDs ...string) *v210.Report {
	t.Helper()

	report := v210.NewReport()
	run := v210.NewRunWithInformationURI(toolName, "https://example.com/"+toolName)

	for _, ruleID := range ruleIDs {
		run.AddRule(ruleID).WithDescription("description of " + ruleID)
		run.CreateResultForRule(ruleID).
			WithLevel("error").
			WithMessage(v210.NewTextMessage(ruleID + " was breached"))
	}

	report.AddRun(run)
	return report
}

func TestV22(t *testing.T) {
	t.Run("appends the runs of every report in order", func(t *testing.T) {
		merged, err := merge.V22(
			newV22Report(t, "tool-a", "rule#1"),
			newV22Report(t, "tool-b", "rule#2"),
			newV22Report(t, "tool-c", "rule#3"),
		)
		require.NoError(t, err)

		require.Len(t, merged.Runs, 3)
		assert.Equal(t, "tool-a", *merged.Runs[0].Tool.Driver.Name)
		assert.Equal(t, "tool-b", *merged.Runs[1].Tool.Driver.Name)
		assert.Equal(t, "tool-c", *merged.Runs[2].Tool.Driver.Name)
	})

	t.Run("keeps every result", func(t *testing.T) {
		merged, err := merge.V22(
			newV22Report(t, "tool-a", "rule#1", "rule#2"),
			newV22Report(t, "tool-b", "rule#3"),
		)
		require.NoError(t, err)

		total := 0
		for _, run := range merged.Runs {
			total += len(run.Results)
		}
		assert.Equal(t, 3, total)
	})

	t.Run("leaves rule indices pointing at the right rules", func(t *testing.T) {
		// Both reports use ruleIndex 0, but for different rules. Appending runs
		// rather than coalescing them is what keeps both correct.
		merged, err := merge.V22(
			newV22Report(t, "tool-a", "rule#1"),
			newV22Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)

		for _, run := range merged.Runs {
			for _, result := range run.Results {
				require.GreaterOrEqual(t, result.RuleIndex, 0)
				require.Less(t, result.RuleIndex, len(run.Tool.Driver.Rules))

				referenced := run.Tool.Driver.Rules[result.RuleIndex]
				assert.Equal(t, *result.RuleID, *referenced.ID,
					"ruleIndex must still resolve to the rule the result names")
			}
		}
	})

	t.Run("leaves artifact indices resolvable", func(t *testing.T) {
		merged, err := merge.V22(
			newV22Report(t, "tool-a", "rule#1"),
			newV22Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)

		for _, run := range merged.Runs {
			require.Len(t, run.Artifacts, 1)
			assert.NotNil(t, run.Artifacts[0].Location.URI)
		}
	})

	t.Run("generates a fresh guid for the merged document", func(t *testing.T) {
		first := newV22Report(t, "tool-a", "rule#1")
		second := newV22Report(t, "tool-b", "rule#2")

		merged, err := merge.V22(first, second)
		require.NoError(t, err)

		assert.NotEmpty(t, merged.Guid)
		assert.NotEqual(t, first.Guid, merged.Guid, "the merge is a new document")
		assert.NotEqual(t, second.Guid, merged.Guid)
	})

	t.Run("unions properties with the earliest report winning", func(t *testing.T) {
		first := newV22Report(t, "tool-a", "rule#1")
		first.Properties.Add("shared", "from-first").Add("only-first", "a").AddTag("alpha")

		second := newV22Report(t, "tool-b", "rule#2")
		second.Properties.Add("shared", "from-second").Add("only-second", "b").AddTag("alpha").AddTag("beta")

		merged, err := merge.V22(first, second)
		require.NoError(t, err)

		assert.Equal(t, "from-first", merged.Properties.Properties["shared"])
		assert.Equal(t, "a", merged.Properties.Properties["only-first"])
		assert.Equal(t, "b", merged.Properties.Properties["only-second"])
		assert.Equal(t, []string{"alpha", "beta"}, merged.Properties.Tags, "tags are deduplicated")
	})

	t.Run("merges inline external properties", func(t *testing.T) {
		first := newV22Report(t, "tool-a", "rule#1")
		first.InlineExternalProperties = []*v22.ExternalProperties{v22.NewExternalProperties()}
		second := newV22Report(t, "tool-b", "rule#2")
		second.InlineExternalProperties = []*v22.ExternalProperties{v22.NewExternalProperties()}

		merged, err := merge.V22(first, second)
		require.NoError(t, err)

		assert.Len(t, merged.InlineExternalProperties, 2)
	})

	t.Run("merging a single report is a valid passthrough", func(t *testing.T) {
		merged, err := merge.V22(newV22Report(t, "tool-a", "rule#1"))
		require.NoError(t, err)

		require.Len(t, merged.Runs, 1)
		assert.NoError(t, merged.Validate())
	})

	t.Run("produces a schema valid report", func(t *testing.T) {
		merged, err := merge.V22(
			newV22Report(t, "tool-a", "rule#1"),
			newV22Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)
		assert.NoError(t, merged.Validate())
	})

	t.Run("rejects zero reports", func(t *testing.T) {
		_, err := merge.V22()
		assert.Error(t, err)
	})

	t.Run("rejects a nil report", func(t *testing.T) {
		_, err := merge.V22(newV22Report(t, "tool-a", "rule#1"), nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "position 1")
	})
}

func TestV210(t *testing.T) {
	t.Run("appends the runs of every report in order", func(t *testing.T) {
		merged, err := merge.V210(
			newV210Report(t, "tool-a", "rule#1"),
			newV210Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)

		require.Len(t, merged.Runs, 2)
		assert.Equal(t, "tool-a", *merged.Runs[0].Tool.Driver.Name)
		assert.Equal(t, "tool-b", *merged.Runs[1].Tool.Driver.Name)
	})

	t.Run("leaves rule indices pointing at the right rules", func(t *testing.T) {
		merged, err := merge.V210(
			newV210Report(t, "tool-a", "rule#1"),
			newV210Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)

		for _, run := range merged.Runs {
			for _, result := range run.Results {
				require.Less(t, result.RuleIndex, len(run.Tool.Driver.Rules))
				assert.Equal(t, *result.RuleID, *run.Tool.Driver.Rules[result.RuleIndex].ID)
			}
		}
	})

	t.Run("unions properties with the earliest report winning", func(t *testing.T) {
		first := newV210Report(t, "tool-a", "rule#1")
		first.Properties.Add("shared", "from-first")

		second := newV210Report(t, "tool-b", "rule#2")
		second.Properties.Add("shared", "from-second").Add("only-second", "b")

		merged, err := merge.V210(first, second)
		require.NoError(t, err)

		assert.Equal(t, "from-first", merged.Properties.Properties["shared"])
		assert.Equal(t, "b", merged.Properties.Properties["only-second"])
	})

	t.Run("produces a schema valid report", func(t *testing.T) {
		merged, err := merge.V210(
			newV210Report(t, "tool-a", "rule#1"),
			newV210Report(t, "tool-b", "rule#2"),
		)
		require.NoError(t, err)
		assert.NoError(t, merged.Validate())
	})

	t.Run("rejects zero reports", func(t *testing.T) {
		_, err := merge.V210()
		assert.Error(t, err)
	})

	t.Run("rejects a nil report", func(t *testing.T) {
		_, err := merge.V210(nil)
		assert.Error(t, err)
	})
}
