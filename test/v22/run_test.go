package test

import "testing"

func Test_new_result_on_run(t *testing.T) {
	given, _, _ := newRunTest(t)

	given.a_result_is_added()
}

func Test_properties_on_a_run(t *testing.T) {
	given, when, then := newRunTest(t)

	expected := `{"tool":{"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[],"properties":{"properties":{"boolean_key":false,"string_key":"string_value"},"tags":[]}}`

	given.properties_added_to_a_run()
	when.the_run_is_json()
	then.string_is_as_expected(expected)
}
