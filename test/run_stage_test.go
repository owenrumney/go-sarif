package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

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
	tool := models.NewTool("tfsec", "https://tfsec.dev")
	rt.run = models.NewRun(tool)
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

	location := &models.Location{
		Uri: locationUri,
	}

	rt.run.GetOrCreateLocation(location)

	return rt
}

func (rt *runTest) and() *runTest {
	return rt
}

func (rt *runTest) the_index_of_location_is(locationUri string, expectedIndex int) *runTest {
	location := &models.Location{
		Uri: locationUri,
	}

	locationIndex, err := rt.run.GetOrCreateLocation(location)
	if err != nil {
		rt.t.Error(err)
	}

	assert.Equal(rt.t, expectedIndex, locationIndex)
	return rt
}

func (rt *runTest) a_result_is_added_to_the_run() *runTest {
	resultLocation := "/tmp/result/code"

	rule := models.NewRule("AWS001", "S3 Bucket has an ACL defined which allows public access.", "https://www.tfsec.dev/docs/aws/AWS001", nil)

	location := &models.PhysicalLocation{
		ArtifactLocation: &models.ArtifactLocation{
			Uri: resultLocation,
		},
		Region: &models.Region{
			StartLine:   1,
			StartColumn: 1,
		},
	}

	result := models.NewResult("error", "Resource 'my_bucket' has an ACL which allows public access.", rule.Id)

	rt.run.AddResultDetails(rule, location, result)
	return rt
}
