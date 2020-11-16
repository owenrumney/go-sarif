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

func (st *sarifTest) contentShouldBe(expected string) {
	assert.Equal(st.t, st.content, expected)
}

func (st *sarifTest) and() *sarifTest {
	return st
}

func (st *sarifTest) aDriverIsAdded() *sarifTest {
	st.sarifReport.AddRun("ESLint", "https://eslint.org")
	return st
}
