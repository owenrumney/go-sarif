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

func (rt *resultTest) aNewResult() {
	rt.result = &models.Result{
		RuleID: "test-rule",
	}

	rt.result.WithLevel("error").
		WithMessage("there was an error")
}

func (rt *resultTest) and() *resultTest {
	return rt
}

func (rt *resultTest) theResultIsDisplayedConvertedAString() {
	jsonContent, err := json.Marshal(rt.result)
	if err != nil {
		rt.t.Error(err)
	}
	rt.content = string(jsonContent)
}

func (rt *resultTest) theResultHasALocationAdded() *resultTest {
	rt.result.WithLocationDetails("/tmp/code/location", 1, 1)
	return rt
}

func (rt *resultTest) theJSONStringRepresentationShouldBe(expected string) {
	assert.Equal(rt.t, expected, rt.content)
}
