package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/owenrumney/go-sarif/models"
)

type resultTest struct {
	t       *testing.T
	result  *models.Result
	content string
}

func createNewResultTest(t *testing.T) (*resultTest, *resultTest, *resultTest) {
	resultTest := &resultTest{
		t: t,
	}
	return resultTest, resultTest, resultTest
}

func (rt *resultTest) a_new_result() {
	rt.result = &models.Result{
		RuleId: "test-rule",
	}

	rt.result.WithLevel("error").
		WithMessage("there was an error")
}

func (rt *resultTest) and() *resultTest {
	return rt
}

func (rt *resultTest) the_result_is_displayed_converted_a_string() {
	jsonContent, err := json.Marshal(rt.result)
	if err != nil {
		rt.t.Error(err)
	}
	rt.content = string(jsonContent)
}

func (rt *resultTest) the_result_has_a_location_added() *resultTest {
	rt.result.WithLocationDetails("/tmp/code/location", 1, 1)
	return rt
}

func (rt *resultTest) the_json_string_representation_should_be(expected string) {
	assert.Equal(rt.t, expected, rt.content)
}
