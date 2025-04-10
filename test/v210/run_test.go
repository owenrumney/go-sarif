package test

import "testing"

func Test_new_result_on_run(t *testing.T) {
	given, _, _ := newRunTest(t)

	given.a_result_is_added()
}

func Test_properties_on_a_run(t *testing.T) {
	given, when, then := newRunTest(t)

	expected := `{"language":"en-US","newlineSequences":["\r\n","\n"],"properties":{"boolean_key":false,"string_key":"string_value"},"tool":{}}`

	given.properties_added_to_a_run()
	when.the_run_is_json()
	then.string_is_as_expected(expected)
}
