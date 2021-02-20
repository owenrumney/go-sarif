package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
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

func Test_create_new_a_new_sarif_file_with_a_driver_in_a_pretty_format(t *testing.T) {
	given, when, then := createNewSarifTest(t)

	expected := `{
  "version": "2.1.0",
  "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationUri": "https://eslint.org"
        }
      }
    }
  ]
}`

	given.aNewSarifReport("2.1.0")
	when.aDriverIsAdded().
		and().theReportIsWrittenToStringInAPrettyFormat()
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

func Test_load_sarif_report_from_string(t *testing.T) {
	given, _, then := createNewSarifTest(t)

	content  := `{
  "version": "2.1.0",
  "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationUri": "https://eslint.org"
        }
      }
    }
  ]
}`

	given.aReportIsLoadedFromString(content)
	then.theReportHasDriverNameAndInformationUri("ESLint","https://eslint.org")
}

func Test_load_sarif_report_from_file(t *testing.T) {
	given, _, then := createNewSarifTest(t)

	content  := `{
  "version": "2.1.0",
  "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationUri": "https://eslint.org"
        }
      }
    }
  ]
}`

	file, err := ioutil.TempFile(os.TempDir(), "sarifReport")
	assert.NoError(t, err)
	defer file.Close()

	ioutil.WriteFile(file.Name(), []byte(content), 755)

	given.aReportIsLoadedFromFile(file.Name())
	then.theReportHasDriverNameAndInformationUri("ESLint","https://eslint.org")
}


func Test_err_on_load_sarif_report_from_file_when_not_exists(t *testing.T) {
	given, _, _ := createNewSarifTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided file path doesn't have a file", err.Error())
		}
	}()
	given.aReportIsLoadedFromFile("")
}

func Test_err_on_load_sarif_report_from_file_when_file_not_legit(t *testing.T) {
	given, _, _ := createNewSarifTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided filepath could not be opened. read /tmp: is a directory", err.Error())
		}
	}()
	given.aReportIsLoadedFromFile("/tmp")
}