package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_new_a_new_empty_sarif_file(t *testing.T) {
	given, when, then := createNewSarifTest(t)

	expected := `{"version":"2.1.0","$schema":"http://json.schemastore.org/sarif-2.1.0-rtm.4","runs":[]}`

	given.aNewSarifReport("2.1.0")
	when.theReportIsWrittenToString()
	then.contentShouldBe(expected)
}

func Test_create_new_a_new_sarif_file_with_a_driver(t *testing.T) {
	given, when, then := createNewSarifTest(t)

	expected := `{"version":"2.1.0","$schema":"http://json.schemastore.org/sarif-2.1.0-rtm.4","runs":[{"tool":{"driver":{"name":"ESLint","informationUri":"https://eslint.org"}}}]}`

	given.aNewSarifReport("2.1.0")
	when.aDriverIsAdded().
		and().theReportIsWrittenToString()
	then.contentShouldBe(expected)
}

func Test_error_when_unsupported_version_requested(t *testing.T) {
	given, _, _ := createNewSarifTest(t)

	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "version [bad_version] is not supported", err.Error())
		}
	}()
	given.aNewSarifReport("bad_version")
}
