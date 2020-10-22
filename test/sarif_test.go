package test

import "testing"

func Test_create_new_sarif_file(t *testing.T) {
	given, when, then := CreateNewSarifHarness(t)
	expected := `{"runs":[]}`
	given.a_new_sarif_report()
	when.the_report_is_written_to_string()
	then.content_should_be(expected)
}
