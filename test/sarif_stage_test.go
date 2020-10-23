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

func (st *sarifTest) a_new_sarif_report(version string) {
	report, err := sarif.New(sarif.Version(version))
	if err != nil {
		panic(err)
	}
	st.sarifReport = report
}

func (st *sarifTest) the_report_is_written_to_string() {
	buf := new(bytes.Buffer)
	err := st.sarifReport.Write(buf)
	if err != nil {
		st.t.Error(err)
	}
	st.content = buf.String()
}

func (st *sarifTest) content_should_be(expected string) {
	assert.Equal(st.t, st.content, expected)
}

func (st *sarifTest) and() *sarifTest {
	return st
}

func (st *sarifTest) a_driver_is_added() *sarifTest {
	st.sarifReport.AddRun("ESLint", "https://eslint.org")
	return st
}
