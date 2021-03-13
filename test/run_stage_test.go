package test

import (
	"testing"

	"github.com/owenrumney/go-sarif/sarif"
)

type runTest struct {
	t   *testing.T
	run *sarif.Run
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
