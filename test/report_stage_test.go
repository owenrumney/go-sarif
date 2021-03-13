package test

import (
	"bytes"
	"testing"

	"github.com/owenrumney/go-sarif/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type reportTest struct {
	t      *testing.T
	report *sarif.Report
}

func newReportTest(t *testing.T) (*reportTest, *reportTest, *reportTest) {
	rt := &reportTest{
		t: t,
	}
	return rt, rt, rt
}

func (r *reportTest) a_new_report() *reportTest {
	report, err := sarif.New(sarif.Version210)
	require.NoError(r.t, err)
	r.report = report
	return r
}

func (r *reportTest) and() *reportTest {
	return r
}

func (r *reportTest) with_a_run_added(tool, informationUri string) *sarif.Run {
	run := sarif.NewRun(tool, informationUri)
	r.report.AddRun(run)
	return run
}

func (r *reportTest) an_artifact_is_added_to_the_run(run *sarif.Run, filename string) *reportTest {
	a := run.AddArtifact()
	a.Location = &sarif.ArtifactLocation{
		URI: &filename,
	}
	return r
}

func (r *reportTest) a_result_is_added_to_the_run(run *sarif.Run, ruleId, messageText string) *reportTest {
	result := run.AddResult(ruleId)
	result.Message = sarif.Message{
		Text: &messageText,
	}
	return r
}

func (r *reportTest) report_text_is(expected string) {
	buffer := bytes.NewBufferString("")
	err := r.report.Write(buffer)
	require.NoError(r.t, err)

	assert.Equal(r.t, expected, buffer.String())
}

func (r *reportTest) a_report_is_loaded_from_a_string(content string) {
	report, err := sarif.FromString(content)
	assert.NoError(r.t, err)
	assert.NotNil(r.t, report)
	r.report = report
}

func (r *reportTest) the_report_has_expected_driver_name_and_information_uri(driverName string, informationUri string) {
	assert.Equal(r.t, driverName, r.report.Runs[0].Tool.Driver.Name)
	assert.Equal(r.t, informationUri, r.report.Runs[0].Tool.Driver.InformationURI)
}

func (r *reportTest) a_report_is_loaded_from_a_file(filename string) {
	report, err := sarif.Open(filename)
	if err != nil {
		panic(err)
	}
	r.report = report

}
