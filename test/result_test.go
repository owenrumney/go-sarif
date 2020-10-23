package test

import (
	"testing"
)

func Test_a_new_result_is_created_as_expected(t *testing.T) {
	given, when, then := createNewResultTest(t)

	expected := `{"level":"error","message":{"text":"there was an error"},"ruleId":"test-rule","ruleIndex":0}`

	given.a_new_result()
	when.the_result_is_displayed_converted_a_string()
	then.the_json_string_representation_should_be(expected)
}

func Test_a_new_result_is_created_with_a_location(t *testing.T) {
	given, when, then := createNewResultTest(t)

	expected := `{"level":"error","message":{"text":"there was an error"},"ruleId":"test-rule","ruleIndex":0,"locations":[{"physicalLocation":{"artifactLocation":{"uri":"/tmp/code/location","index":0},"region":{"startLine":1,"startColumn":1}}}]}`

	given.a_new_result()
	when.the_result_has_a_location_added().
		and().
		the_result_is_displayed_converted_a_string()
	then.the_json_string_representation_should_be(expected)

}
