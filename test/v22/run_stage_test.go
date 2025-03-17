package test

import (
	"encoding/json"
	"testing"

	"github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type runTest struct {
	t          *testing.T
	run        *sarif.Run
	jsonString string
}

func newRunTest(t *testing.T) (*runTest, *runTest, *runTest) {
	r := &runTest{
		t:   t,
		run: sarif.NewRun().WithTool(sarif.NewTool()),
	}
	return r, r, r
}

func (r *runTest) a_result_is_added() {
	r.run.AddResult(sarif.NewResult().WithRuleID("rule_id"))
}

func (r *runTest) properties_added_to_a_run() {
	pb := sarif.NewPropertyBag()

	pb.Add("string_key", "string_value")
	pb.Add("boolean_key", false)

	r.run.WithProperties(pb)
}

func (r *runTest) the_run_is_json() {
	b, err := json.Marshal(r.run)
	require.NoError(r.t, err)
	r.jsonString = string(b)
}

func (r *runTest) string_is_as_expected(expected string) {
	assert.Equal(r.t, expected, r.jsonString)
}
