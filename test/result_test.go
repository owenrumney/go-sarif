package test

import (
	"testing"
)

func Test_a_new_result_is_created_as_expected(t *testing.T) {
	given, when, then := createNewResultTest(t)

	expected := `{"level":"error","message":{"text":"there was an error"},"ruleId":"test-rule","ruleIndex":0}`

	given.aNewResult()
	when.theResultIsDisplayedConvertedAString()
	then.theJSONStringRepresentationShouldBe(expected)
}

func Test_a_new_result_is_created_with_a_location(t *testing.T) {
	given, when, then := createNewResultTest(t)

	expected := `{"level":"error","message":{"text":"there was an error"},"ruleId":"test-rule","ruleIndex":0,"locations":[{"physicalLocation":{"artifactLocation":{"uri":"/tmp/code/location","index":0},"region":{"startLine":1,"startColumn":1}}}]}`

	given.aNewResult()
	when.theResultHasALocationAdded().
		and().
		theResultIsDisplayedConvertedAString()
	then.theJSONStringRepresentationShouldBe(expected)

}
