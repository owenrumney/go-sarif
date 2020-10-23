package test

import (
	"testing"
)

func Test_create_new_empty_run_looks_as_expected(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}}}`

	given.a_new_run_is_created()
	when.the_run_is_converted_to_a_string()
	then.the_json_string_representation_of_the_run_should_be(expected)
}

func Test_create_run_with_artifact_as_expected(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"artifacts":[{"location":{"uri":"/tmp/code/location"}}]}`

	given.a_new_run_is_created()
	when.an_artifact_is_added_to_the_run("/tmp/code/location").
		and().
		the_run_is_converted_to_a_string()
	then.the_json_string_representation_of_the_run_should_be(expected)
}

func Test_getting_the_location_index_for_an_existing_run(t *testing.T) {
	given, when, then := createNewRunTest(t)

	given.a_new_run_is_created()
	when.an_artifact_is_added_to_the_run("/tmp/location/1").
		and().
		an_artifact_is_added_to_the_run("/tmp/location/2")
	then.the_index_of_location_is("/tmp/location/1", 0).
		and().
		the_index_of_location_is("/tmp/location/2", 1)

}

func Test_create_a_run_with_a_result_added(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev","rules":[{"id":"AWS001","shortDescription":{"text":"S3 Bucket has an ACL defined which allows public access."},"helpUri":"https://www.tfsec.dev/docs/aws/AWS001","properties":{"propertyName":"propertyValue"}}]}},"artifacts":[{"location":{"uri":"/tmp/result/code"}}],"results":[{"level":"error","message":{"text":"Resource 'my_bucket' has an ACL which allows public access."},"ruleId":"AWS001","ruleIndex":0,"locations":[{"physicalLocation":{"artifactLocation":{"uri":"/tmp/result/code","index":0},"region":{"startLine":1,"startColumn":1}}}]}]}`

	given.a_new_run_is_created()
	when.a_result_is_added_to_the_run().and().
		the_run_is_converted_to_a_string()
	then.the_json_string_representation_of_the_run_should_be(expected)
}
