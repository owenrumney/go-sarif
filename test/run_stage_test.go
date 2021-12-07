package test

import (
	"encoding/json"
	"testing"

	"github.com/owenrumney/go-sarif/sarif"
	"github.com/stretchr/testify/assert"
)

type runTest struct {
	t          *testing.T
	run        *sarif.Run
	jsonString string
}

func newRunTest(t *testing.T) (*runTest, *runTest, *runTest) {
	r := &runTest{
		t: t,
		run: &sarif.Run{
			Tool:      sarif.Tool{},
			Artifacts: []*sarif.Artifact{},
			Results:   []*sarif.Result{},
		},
	}
	return r, r, r
}

func (r *runTest) a_result_is_added() {
	r.run.AddResult("rule1")
}

func (r *runTest) properties_added_to_a_run() {
	pb := sarif.NewPropertyBag()

	pb.AddString("string_key", "string_value")
	pb.AddBoolean("boolean_key", false)

	r.run.AttachPropertyBag(pb)
}

func (r *runTest) the_run_is_json() {
	b, err := json.Marshal(r.run)
	assert.Nil(r.t, err)
	r.jsonString = string(b)
}

func (r *runTest) string_is_as_expected(expected string) {
	assert.Equal(r.t, expected, r.jsonString)
}
