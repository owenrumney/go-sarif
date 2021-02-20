package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/owenrumney/go-sarif/sarif"
)

type sarifTest struct {
	t           *testing.T
	sarifReport *sarif.Report
	content     string
}

func createNewSarifTest(t *testing.T) (*sarifTest, *sarifTest, *sarifTest) {
	sarifTest := &sarifTest{
		t: t,
	}
	return sarifTest, sarifTest, sarifTest
}

func (st *sarifTest) aNewSarifReport(version string) {
	report, err := sarif.New(sarif.Version(version))
	if err != nil {
		panic(err)
	}
	st.sarifReport = report
}

func (st *sarifTest) theReportIsWrittenToString() {
	buf := new(bytes.Buffer)
	err := st.sarifReport.Write(buf)
	if err != nil {
		st.t.Error(err)
	}
	st.content = buf.String()
}


func (st *sarifTest) theReportIsWrittenToStringInAPrettyFormat() {
	buf := new(bytes.Buffer)
	err := st.sarifReport.PrettyWrite(buf)
	if err != nil {
		st.t.Error(err)
	}
	st.content = buf.String()
}

func (st *sarifTest) contentShouldBe(expected string) {
	assert.Equal(st.t, expected, st.content)
}

func (st *sarifTest) and() *sarifTest {
	return st
}

func (st *sarifTest) aDriverIsAdded() *sarifTest {
	st.sarifReport.AddRun("ESLint", "https://eslint.org")
	return st
}

func (st *sarifTest) aReportIsLoadedFromString(content string) {
	report, err := sarif.FromString(content)
	assert.NoError(st.t, err)
	assert.NotNil(st.t, report)
	st.sarifReport = report
}

func (st *sarifTest) theReportHasDriverNameAndInformationUri(driverName string, informationUri string) {
	assert.Equal(st.t, driverName, st.sarifReport.Runs[0].Tool.Driver.Name)
	assert.Equal(st.t, informationUri, st.sarifReport.Runs[0].Tool.Driver.InformationURI)
}

func (st *sarifTest) aReportIsLoadedFromFile(filename string) {
	report, err := sarif.Open(filename)
	if err != nil {
		panic(err)
	}
	st.sarifReport = report

}