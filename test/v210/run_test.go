package test

import "testing"

func Test_new_result_on_run(t *testing.T) {
	given, _, _ := newRunTest(t)

	given.a_result_is_added()
}

func Test_properties_on_a_run(t *testing.T) {
	given, when, then := newRunTest(t)

	expected := `{"taxonomies":[],"invocations":[],"addresses":[],"properties":{"properties":{"boolean_key":false,"string_key":"string_value"},"tags":[]},"artifacts":[],"threadFlowLocations":[],"webRequests":[],"webResponses":[],"versionControlProvenance":[],"results":[],"runAggregates":[],"newlineSequences":["\r\n","\n"],"translations":[],"policies":[],"tool":{"extensions":[]},"logicalLocations":[],"graphs":[],"redactionTokens":[]}`

	given.properties_added_to_a_run()
	when.the_run_is_json()
	then.string_is_as_expected(expected)
}
