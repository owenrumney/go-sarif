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

func (rt *runTest) aNewRunIsCreated() {
	rt.run = models.NewRun("tfsec", "https://tfsec.dev")
}

func (rt *runTest) theRunIsConvertedToAString() {
	jsonContent, err := json.Marshal(rt.run)
	if err != nil {
		rt.t.Error(err)
	}
	rt.content = string(jsonContent)
}

func (rt *runTest) theJSONStringRepresentationOfTheRunShouldBe(expected string) {
	assert.Equal(rt.t, expected, rt.content)
}

func (rt *runTest) anArtifactIsAddedToTheRun(locationURI string) *runTest {
	rt.run.AddArtifact(locationURI)
	return rt
}

func (rt *runTest) and() *runTest {
	return rt
}

func (rt *runTest) theIndexOfLocationIs(locationURI string, expectedIndex int) *runTest {
	locationIndex := rt.run.AddArtifact(locationURI)
	assert.Equal(rt.t, expectedIndex, locationIndex)
	return rt
}

func (rt *runTest) aResultIsAddedToTheRun() *runTest {
	resultLocation := "/tmp/result/code"

	rule := rt.run.AddRule("AWS001").
		WithDescription("S3 Bucket has an ACL defined which allows public access.").
		WithHelpURI("https://www.tfsec.dev/docs/aws/AWS001").
		WithProperties(map[string]string{"propertyName": "propertyValue"})

	result := rt.run.AddResult(rule.ID).
		WithLevel("error").
		WithMessage("Resource 'my_bucket' has an ACL which allows public access.").
		WithLocationDetails(resultLocation, 1, 1)

	rt.run.AddResultDetails(rule, result, resultLocation)
	return rt
}

func (rt *runTest) aResultIsAddedToTheRunWithHelpText() *runTest {
	resultLocation := "/tmp/result/code"

	rule := rt.run.AddRule("AWS001").
		WithDescription("S3 Bucket has an ACL defined which allows public access.").
		WithHelp("you can learn more about this check https://www.tfsec.dev/docs/aws/AWS001").
		WithProperties(map[string]string{"propertyName": "propertyValue"})

	result := rt.run.AddResult(rule.ID).
		WithLevel("error").
		WithMessage("Resource 'my_bucket' has an ACL which allows public access.").
		WithLocationDetails(resultLocation, 1, 1)

	rt.run.AddResultDetails(rule, result, resultLocation)
	return rt
}

func (rt *runTest) gettingRuleByIdReturnsRule() {
	rule, err := rt.run.GetRuleById("AWS001")

	assert.NoError(rt.t, err)
	assert.Equal(rt.t, "AWS001", rule.ID)
	assert.Equal(rt.t, "S3 Bucket has an ACL defined which allows public access.", rule.ShortDescription.Text)
}
