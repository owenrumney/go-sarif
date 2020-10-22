package models

type Result struct {
	Level     string              `json:"level"`
	Message   *TextBlock          `json:"message"`
	RuleId    string              `json:"ruleId"`
	RuleIndex int                 `json:"ruleIndex"`
	Locations []*PhysicalLocation `json:"locations"`
}

type PhysicalLocation struct {
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`
	Region           *Region           `json:"region"`
}

type Region struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}

type ArtifactLocation struct {
	Uri   string `json:"uri"`
	Index int    `json:"index,omitempty"`
}

type Location struct {
	Uri string `json:"uri"`
}

func CreateResult(level, message, ruleId string) *Result {
	return &Result{
		Level: level,
		Message: &TextBlock{
			Text: message,
		},
		RuleId: ruleId,
	}
}

func (r Result) AddLocation(location *PhysicalLocation) {
	r.Locations = append(r.Locations, location)
}
