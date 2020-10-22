package models

type Run struct {
	Tool      *Tool       `json:"tool"`
	Artifacts []*Location `json:"artifacts"`
	Results   []*Result   `json:"results"`
}

func CreateRun(tool *Tool) *Run {
	return &Run{
		Tool: tool,
	}
}

func (run *Run) getOrCreate(location *ArtifactLocation) (int, error) {
	for i, l := range run.Artifacts {
		if l.Uri == location.Uri {
			return i, nil
		}
	}
	run.Artifacts = append(run.Artifacts, &Location{
		Uri: location.Uri,
	})
	return len(run.Artifacts), nil
}

func (run *Run) addResult(result *Result) {
	run.Results = append(run.Results, result)
}

func (run *Run) AddEntry(rule *Rule, location *PhysicalLocation, result *Result) error {
	ruleIndex, err := run.Tool.Driver.GetOrCreate(rule)
	if err != nil {
		return err
	}
	result.RuleIndex = ruleIndex
	locationIndex, err := run.getOrCreate(location.ArtifactLocation)
	if err != nil {
		return nil
	}
	location.ArtifactLocation.Index = locationIndex
	result.AddLocation(location)
	return nil
}
