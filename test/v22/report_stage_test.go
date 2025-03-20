package test

import (
	"bytes"
	"testing"

	"github.com/owenrumney/go-sarif/v3/pkg/report"
	"github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
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
	report := report.NewV22Report()
	r.report = report
	return r
}

func (r *reportTest) and() *reportTest {
	return r
}

func (r *reportTest) with_a_run_added(tool, informationURI string) *sarif.Run {
	run := sarif.NewRunWithInformationURI(tool, informationURI)
	r.report.AddRun(run)
	return run
}

func (r *reportTest) with_a_run_with_empty_result_added(tool, informationURI string) *sarif.Run {
	run := sarif.NewRunWithInformationURI(tool, informationURI)
	r.report.AddRun(run)
	return run
}

func (r *reportTest) an_artifact_is_added_to_the_run(run *sarif.Run, filename string) *reportTest {
	a := sarif.NewArtifact().WithLocation(
		sarif.NewArtifactLocation().WithURI(filename),
	)
	run.AddArtifact(a)
	return r
}

func (r *reportTest) some_properties_are_added_to_the_run(run *sarif.Run) *reportTest {
	pb := sarif.NewPropertyBag()
	pb.Add("string_property", "this is a string")
	pb.Add("integer_property", 10)
	run.WithProperties(pb)
	return r
}

func (r *reportTest) report_text_is(expected string) {
	buffer := bytes.NewBufferString("")
	r.report.Guid = "516d8cc4-18b1-463e-9e0e-417473149927"
	err := r.report.Write(buffer)
	require.NoError(r.t, err)

	assert.JSONEq(r.t, expected, buffer.String())
}

func (r *reportTest) a_report_is_loaded_from_a_string(content string) {
	report, err := sarif.FromString(content)
	require.NoError(r.t, err)
	assert.NotNil(r.t, report)
	r.report = report
}

func (r *reportTest) the_report_has_expected_driver_name_and_information_uri(driverName, informationURI string) {
	assert.Equal(r.t, driverName, r.report.Runs[0].Tool.Driver.Name)
	assert.Equal(r.t, informationURI, r.report.Runs[0].Tool.Driver.InformationURI)
}

func (r *reportTest) a_report_is_loaded_from_a_file(filename string) {
	report, err := sarif.Open(filename)
	if err != nil {
		panic(err)
	}
	r.report = report

}
