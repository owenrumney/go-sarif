package test

import "testing"

func Test_create_new_a_new_empty_sarif_file(t *testing.T) {
	given, when, then := CreateNewSarifHarness(t)

	expected := `{"version":"2.1.0","$schema":"http://json.schemastore.org/sarif-2.1.0-rtm.4","runs":[]}`

	given.a_new_sarif_report()
	when.the_report_is_written_to_string()
	then.content_should_be(expected)
}

func Test_create_new_a_new_sarif_file_with_a_driver(t *testing.T) {
	given, when, then := CreateNewSarifHarness(t)

	expected := `{"version":"2.1.0","$schema":"http://json.schemastore.org/sarif-2.1.0-rtm.4","runs":[{"tool":{"driver":{"name":"ESLint","informationUri":"https://eslint.org"}}}]}`

	given.a_new_sarif_report()
	when.a_driver_is_added().
		and().the_report_is_written_to_string()
	then.content_should_be(expected)
}
