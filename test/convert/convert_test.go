package convert_test

import (
	"testing"

	"github.com/owenrumney/go-sarif/v3/pkg/report/convert"
	v210 "github.com/owenrumney/go-sarif/v3/pkg/report/v210/sarif"
	v22 "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newV210Report(t *testing.T) *v210.Report {
	t.Helper()

	report := v210.NewReport()
	run := v210.NewRunWithInformationURI("test-tool", "https://example.com")

	run.AddRule("rule#1").
		WithDescription("a rule that matters").
		WithHelpURI("https://example.com/rules/1")

	run.AddDistinctArtifact("file:///code/main.tf")

	run.CreateResultForRule("rule#1").
		WithLevel("error").
		WithMessage(v210.NewTextMessage("the rule was breached")).
		AddLocation(v210.NewLocationWithPhysicalLocation(
			v210.NewPhysicalLocation().
				WithArtifactLocation(v210.NewSimpleArtifactLocation("file:///code/main.tf")).
				WithRegion(v210.NewSimpleRegion(1, 4)),
		))

	report.AddRun(run)
	return report
}

func newV22Report(t *testing.T) *v22.Report {
	t.Helper()

	report := v22.NewReport()
	run := v22.NewRunWithInformationURI("test-tool", "https://example.com")

	run.AddRule("rule#1").
		WithDescription("a rule that matters").
		WithHelpURI("https://example.com/rules/1")

	run.AddDistinctArtifact("file:///code/main.tf")

	run.CreateResultForRule("rule#1").
		WithLevel("error").
		WithMessage(v22.NewTextMessage("the rule was breached")).
		AddLocation(v22.NewLocationWithPhysicalLocation(
			v22.NewPhysicalLocation().
				WithArtifactLocation(v22.NewSimpleArtifactLocation("file:///code/main.tf")).
				WithRegion(v22.NewSimpleRegion(1, 4)),
		))

	report.AddRun(run)
	return report
}

func TestToV22(t *testing.T) {
	t.Run("upgrades version and schema", func(t *testing.T) {
		converted, err := convert.ToV22(newV210Report(t))
		require.NoError(t, err)

		assert.Equal(t, "2.2", converted.Version)
		assert.Contains(t, converted.Schema, "sarif-2-2")
	})

	t.Run("generates a report guid", func(t *testing.T) {
		converted, err := convert.ToV22(newV210Report(t))
		require.NoError(t, err)

		assert.NotEmpty(t, converted.Guid, "2.2 requires a report level guid")
	})

	t.Run("carries the run content across", func(t *testing.T) {
		converted, err := convert.ToV22(newV210Report(t))
		require.NoError(t, err)

		require.Len(t, converted.Runs, 1)
		run := converted.Runs[0]

		require.NotNil(t, run.Tool)
		require.NotNil(t, run.Tool.Driver)
		assert.Equal(t, "test-tool", *run.Tool.Driver.Name)

		require.Len(t, run.Tool.Driver.Rules, 1)
		assert.Equal(t, "rule#1", *run.Tool.Driver.Rules[0].ID)

		require.Len(t, run.Results, 1)
		assert.Equal(t, "error", run.Results[0].Level)
		assert.Equal(t, "the rule was breached", *run.Results[0].Message.Text)

		require.Len(t, run.Artifacts, 1)
		assert.Equal(t, "file:///code/main.tf", *run.Artifacts[0].Location.URI)
	})

	t.Run("is lossless even under strict conversion", func(t *testing.T) {
		converted, err := convert.ToV22(newV210Report(t), convert.WithStrictConversion())
		require.NoError(t, err, "upgrading should never drop data")
		assert.NotNil(t, converted)
	})

	t.Run("produces a schema valid report", func(t *testing.T) {
		converted, err := convert.ToV22(newV210Report(t))
		require.NoError(t, err)
		assert.NoError(t, converted.Validate())
	})

	t.Run("rejects a nil report", func(t *testing.T) {
		_, err := convert.ToV22(nil)
		assert.Error(t, err)
	})
}

func TestToV210(t *testing.T) {
	t.Run("downgrades version and schema", func(t *testing.T) {
		converted, err := convert.ToV210(newV22Report(t))
		require.NoError(t, err)

		assert.Equal(t, "2.1.0", converted.Version)
		assert.Contains(t, converted.Schema, "2.1.0")
	})

	t.Run("carries the run content across", func(t *testing.T) {
		converted, err := convert.ToV210(newV22Report(t))
		require.NoError(t, err)

		require.Len(t, converted.Runs, 1)
		run := converted.Runs[0]

		require.NotNil(t, run.Tool)
		require.NotNil(t, run.Tool.Driver)
		assert.Equal(t, "test-tool", *run.Tool.Driver.Name)

		require.Len(t, run.Results, 1)
		assert.Equal(t, "error", run.Results[0].Level)
		assert.Equal(t, "the rule was breached", *run.Results[0].Message.Text)
	})

	t.Run("reports the dropped report guid", func(t *testing.T) {
		var losses []convert.Loss
		_, err := convert.ToV210(newV22Report(t), convert.WithLossHandler(func(l convert.Loss) {
			losses = append(losses, l)
		}))
		require.NoError(t, err)

		require.NotEmpty(t, losses, "2.1.0 has no report level guid")
		assert.Equal(t, "guid", losses[0].Path)
	})

	t.Run("drops relatedLocations from notifications", func(t *testing.T) {
		report := newV22Report(t)
		notification := v22.NewNotification()
		notification.RelatedLocations = []*v22.Location{
			v22.NewLocationWithPhysicalLocation(v22.NewPhysicalLocation().
				WithArtifactLocation(v22.NewSimpleArtifactLocation("file:///code/main.tf"))),
		}
		invocation := v22.NewInvocation()
		invocation.ToolExecutionNotifications = []*v22.Notification{notification}
		report.Runs[0].Invocations = []*v22.Invocation{invocation}

		var paths []string
		converted, err := convert.ToV210(report, convert.WithLossHandler(func(l convert.Loss) {
			paths = append(paths, l.Path)
		}))
		require.NoError(t, err)

		assert.Contains(t, paths, "runs[0].invocations[0].toolExecutionNotifications[0].relatedLocations")
		require.Len(t, converted.Runs[0].Invocations, 1, "the invocation itself must survive")
	})

	t.Run("strict conversion fails on loss", func(t *testing.T) {
		_, err := convert.ToV210(newV22Report(t), convert.WithStrictConversion())
		require.Error(t, err)

		var lossErr *convert.LossyConversionError
		require.ErrorAs(t, err, &lossErr)
		assert.NotEmpty(t, lossErr.Losses)
	})

	t.Run("strict conversion succeeds when nothing is lost", func(t *testing.T) {
		report := newV22Report(t)
		report.Guid = "" // the only 2.2 only field this report populates

		converted, err := convert.ToV210(report, convert.WithStrictConversion())
		require.NoError(t, err)
		assert.Equal(t, "2.1.0", converted.Version)
	})

	t.Run("defaults the language 2.2 leaves unset", func(t *testing.T) {
		source := newV22Report(t)
		require.Nil(t, source.Runs[0].Language, "2.2 leaves language unset")

		converted, err := convert.ToV210(source)
		require.NoError(t, err)

		assert.Equal(t, "en-US", converted.Runs[0].Language)
		assert.Equal(t, "en-US", converted.Runs[0].Tool.Driver.Language)
	})

	t.Run("keeps a language that 2.2 did set", func(t *testing.T) {
		source := newV22Report(t)
		language := v22.Language("fr-FR")
		source.Runs[0].Language = &language

		converted, err := convert.ToV210(source)
		require.NoError(t, err)

		assert.Equal(t, "fr-FR", converted.Runs[0].Language)
	})

	t.Run("produces a schema valid report", func(t *testing.T) {
		converted, err := convert.ToV210(newV22Report(t))
		require.NoError(t, err)
		assert.NoError(t, converted.Validate())
	})

	t.Run("rejects a nil report", func(t *testing.T) {
		_, err := convert.ToV210(nil)
		assert.Error(t, err)
	})
}

func TestRoundTrip(t *testing.T) {
	t.Run("v210 survives a round trip unchanged", func(t *testing.T) {
		original := newV210Report(t)

		upgraded, err := convert.ToV22(original)
		require.NoError(t, err)

		roundTripped, err := convert.ToV210(upgraded)
		require.NoError(t, err)

		assert.Equal(t, original.Version, roundTripped.Version)
		assert.Equal(t, original.Schema, roundTripped.Schema)
		require.Len(t, roundTripped.Runs, len(original.Runs))

		originalRun, roundTrippedRun := original.Runs[0], roundTripped.Runs[0]
		assert.Equal(t, *originalRun.Tool.Driver.Name, *roundTrippedRun.Tool.Driver.Name)
		assert.Len(t, roundTrippedRun.Results, len(originalRun.Results))
		assert.Equal(t, originalRun.Results[0].Level, roundTrippedRun.Results[0].Level)
		assert.Equal(t, *originalRun.Results[0].Message.Text, *roundTrippedRun.Results[0].Message.Text)
	})

	t.Run("numeric values keep their precision", func(t *testing.T) {
		original := newV210Report(t)
		original.Runs[0].Results[0].Locations[0].PhysicalLocation.Region.WithByteOffset(9007199254740993)

		upgraded, err := convert.ToV22(original)
		require.NoError(t, err)

		region := upgraded.Runs[0].Results[0].Locations[0].PhysicalLocation.Region
		assert.Equal(t, 9007199254740993, region.ByteOffset)
	})
}
