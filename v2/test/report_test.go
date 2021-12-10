package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_new_simple_report_with_single_run(t *testing.T) {
	given, _, then := newReportTest(t)

	given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	then.report_text_is(`{"version":"2.1.0","$schema":"https://json.schemastore.org/sarif-2.1.0-rtm.5.json","runs":[{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"results":[]}]}`)
}

func Test_new_report_with_empty_run(t *testing.T) {
	given, _, then := newReportTest(t)

	given.a_new_report().
		with_a_run_with_empty_result_added("tfsec", "https://tfsec.dev")
	then.report_text_is(`{"version":"2.1.0","$schema":"https://json.schemastore.org/sarif-2.1.0-rtm.5.json","runs":[{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"results":[]}]}`)
}

func Test_new_simple_report_with_artifact(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.an_artifact_is_added_to_the_run(run, "file://broken.go")
	then.report_text_is(`{"version":"2.1.0","$schema":"https://json.schemastore.org/sarif-2.1.0-rtm.5.json","runs":[{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"artifacts":[{"location":{"uri":"file://broken.go"},"length":-1}],"results":[]}]}`)
}

func Test_new_simple_report_with_propertybag(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.some_properties_are_added_to_the_run(run)
	then.report_text_is(`{"version":"2.1.0","$schema":"https://json.schemastore.org/sarif-2.1.0-rtm.5.json","runs":[{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"results":[],"properties":{"integer_property":10,"string_property":"this is a string"}}]}`)
}

func Test_new_simple_report_with_duplicate_artifact(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.an_artifact_is_added_to_the_run(run, "file://broken.go").
		and().
		an_artifact_is_added_to_the_run(run, "file://broken.go")
	then.report_text_is(`{"version":"2.1.0","$schema":"https://json.schemastore.org/sarif-2.1.0-rtm.5.json","runs":[{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"artifacts":[{"location":{"uri":"file://broken.go"},"length":-1},{"location":{"uri":"file://broken.go"},"length":-1}],"results":[]}]}`)
}

func Test_load_sarif_from_string(t *testing.T) {
	given, _, then := newReportTest(t)

	content := `{
  "version": "2.1.0",
  "$schema": "https://json.schemastore.org/sarif-2.1.0-rtm.5.json",
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

	given.a_report_is_loaded_from_a_string(content)
	then.the_report_has_expected_driver_name_and_information_uri("ESLint", "https://eslint.org")
}

func Test_load_sarif_report_from_file(t *testing.T) {
	given, _, then := newReportTest(t)

	content := `{
  "version": "2.1.0",
  "$schema": "https://json.schemastore.org/sarif-2.1.0-rtm.5.json",
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

	given.a_report_is_loaded_from_a_file(file.Name())
	then.the_report_has_expected_driver_name_and_information_uri("ESLint", "https://eslint.org")
}

func Test_err_on_load_sarif_report_from_file_when_not_exists(t *testing.T) {
	given, _, _ := newReportTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided file path doesn't have a file", err.Error())
		}
	}()
	given.a_report_is_loaded_from_a_file("")
}

func Test_err_on_load_sarif_report_from_file_when_file_not_legit(t *testing.T) {
	given, _, _ := newReportTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided filepath could not be opened. read /tmp: is a directory", err.Error())
		}
	}()
	given.a_report_is_loaded_from_a_file("/tmp")
}
