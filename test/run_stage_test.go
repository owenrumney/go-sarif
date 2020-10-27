package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/owenrumney/go-sarif/models"
)

type runTest struct {
	t       *testing.T
	run     *models.Run
	content string
}

func createNewRunTest(t *testing.T) (*runTest, *runTest, *runTest) {
	runTest := &runTest{
		t: t,
	}
	return runTest, runTest, runTest
}

func (rt *runTest) a_new_run_is_created() {
	rt.run = models.NewRun("tfsec", "https://tfsec.dev")
}

func (rt *runTest) the_run_is_converted_to_a_string() {
	jsonContent, err := json.Marshal(rt.run)
	if err != nil {
		rt.t.Error(err)
	}
	rt.content = string(jsonContent)
}

func (rt *runTest) the_json_string_representation_of_the_run_should_be(expected string) {
	assert.Equal(rt.t, expected, rt.content)
}

func (rt *runTest) an_artifact_is_added_to_the_run(locationUri string) *runTest {
	rt.run.AddArtifact(locationUri)
	return rt
}

func (rt *runTest) and() *runTest {
	return rt
}

func (rt *runTest) the_index_of_location_is(locationUri string, expectedIndex int) *runTest {
	locationIndex := rt.run.AddArtifact(locationUri)
	assert.Equal(rt.t, expectedIndex, locationIndex)
	return rt
}

func (rt *runTest) a_result_is_added_to_the_run() *runTest {
	resultLocation := "/tmp/result/code"

	rule := rt.run.AddRule("AWS001").
		WithDescription("S3 Bucket has an ACL defined which allows public access.").
		WithHelpUri("https://www.tfsec.dev/docs/aws/AWS001").
		WithProperties(map[string]string{"propertyName": "propertyValue"})

	result := rt.run.AddResult(rule.Id).
		WithLevel("error").
		WithMessage("Resource 'my_bucket' has an ACL which allows public access.").
		WithLocationDetails(resultLocation, 1, 1)

	rt.run.AddResultDetails(rule, result, resultLocation)
	return rt
}

func (rt *runTest) a_result_is_added_to_the_run_with_help_text() *runTest {
	resultLocation := "/tmp/result/code"

	rule := rt.run.AddRule("AWS001").
		WithDescription("S3 Bucket has an ACL defined which allows public access.").
		WithHelp("you can learn more about this check https://www.tfsec.dev/docs/aws/AWS001").
		WithProperties(map[string]string{"propertyName": "propertyValue"})

	result := rt.run.AddResult(rule.Id).
		WithLevel("error").
		WithMessage("Resource 'my_bucket' has an ACL which allows public access.").
		WithLocationDetails(resultLocation, 1, 1)

	rt.run.AddResultDetails(rule, result, resultLocation)
	return rt
}
