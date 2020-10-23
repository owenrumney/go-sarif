package models

// Result represents the results block in the sarif report
type Result struct {
	Level     string            `json:"level"`
	Message   *TextBlock        `json:"message"`
	RuleId    string            `json:"ruleId"`
	RuleIndex int               `json:"ruleIndex"`
	Locations []*ResultLocation `json:"locations,omitempty"`
}

type ResultLocation struct {
	PhysicalLocation *PhysicalLocation `json:"physicalLocation,omitempty"`
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
	Index int    `json:"index"`
}

type Location struct {
	Uri string `json:"uri"`
}

func NewResult(level, message, ruleId string) *Result {
	return &Result{
		Level: level,
		Message: &TextBlock{
			Text: message,
		},
		RuleId: ruleId,
	}
}

func (r *Result) AddLocation(location *PhysicalLocation) {
	r.Locations = append(r.Locations, &ResultLocation{
		PhysicalLocation: location,
	})
}
