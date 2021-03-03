package test

import (
	"testing"
)

func Test_create_new_empty_run_looks_as_expected(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}}}`

	given.aNewRunIsCreated()
	when.theRunIsConvertedToAString()
	then.theJSONStringRepresentationOfTheRunShouldBe(expected)
}

func Test_create_run_with_artifact_as_expected(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev"}},"artifacts":[{"location":{"uri":"/tmp/code/location"},"length":0}]}`

	given.aNewRunIsCreated()
	when.anArtifactIsAddedToTheRun("/tmp/code/location").
		and().
		theRunIsConvertedToAString()
	then.theJSONStringRepresentationOfTheRunShouldBe(expected)
}

func Test_getting_the_location_index_for_an_existing_run(t *testing.T) {
	given, when, then := createNewRunTest(t)

	given.aNewRunIsCreated()
	when.anArtifactIsAddedToTheRun("/tmp/location/1").
		and().
		anArtifactIsAddedToTheRun("/tmp/location/2")
	then.theIndexOfLocationIs("/tmp/location/1", 0).
		and().
		theIndexOfLocationIs("/tmp/location/2", 1)

}

func Test_create_a_run_with_a_result_added(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev","rules":[{"id":"AWS001","shortDescription":{"text":"S3 Bucket has an ACL defined which allows public access."},"helpUri":"https://www.tfsec.dev/docs/aws/AWS001","properties":{"propertyName":"propertyValue"}}]}},"artifacts":[{"location":{"uri":"/tmp/result/code"},"length":0}],"results":[{"ruleId":"AWS001","ruleIndex":0,"level":"error","message":{"text":"Resource 'my_bucket' has an ACL which allows public access."},"locations":[{"physicalLocation":{"artifactLocation":{"uri":"/tmp/result/code","index":0},"region":{"startLine":1,"startColumn":1}}}]}]}`

	given.aNewRunIsCreated()
	when.aResultIsAddedToTheRun().and().
		theRunIsConvertedToAString()
	then.theJSONStringRepresentationOfTheRunShouldBe(expected)
}

func Test_create_a_run_with_a_result_added_and_help_text_provided(t *testing.T) {
	given, when, then := createNewRunTest(t)

	expected := `{"tool":{"driver":{"name":"tfsec","informationUri":"https://tfsec.dev","rules":[{"id":"AWS001","shortDescription":{"text":"S3 Bucket has an ACL defined which allows public access."},"help":{"text":"you can learn more about this check https://www.tfsec.dev/docs/aws/AWS001"},"properties":{"propertyName":"propertyValue"}}]}},"artifacts":[{"location":{"uri":"/tmp/result/code"},"length":0}],"results":[{"ruleId":"AWS001","ruleIndex":0,"level":"error","message":{"text":"Resource 'my_bucket' has an ACL which allows public access."},"locations":[{"physicalLocation":{"artifactLocation":{"uri":"/tmp/result/code","index":0},"region":{"startLine":1,"startColumn":1}}}]}]}`

	given.aNewRunIsCreated()
	when.aResultIsAddedToTheRunWithHelpText().and().
		theRunIsConvertedToAString()
	then.theJSONStringRepresentationOfTheRunShouldBe(expected)
}

func Test_create_a_run_with_a_rule_and_get_the_rule_by_id(t *testing.T) {
	given, when, then := createNewRunTest(t)

	given.aNewRunIsCreated()
	when.aResultIsAddedToTheRunWithHelpText()
	then.gettingRuleByIdReturnsRule()
}
