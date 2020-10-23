package models

type Run struct {
	Tool      *Tool              `json:"tool"`
	Artifacts []*LocationWrapper `json:"artifacts,omitempty"`
	Results   []*Result          `json:"results,omitempty"`
}

type LocationWrapper struct {
	Location *Location `json:"location,omitentry"`
}

// NewRun will create a new Run
func NewRun(tool *Tool) *Run {
	return &Run{
		Tool: tool,
	}
}

func (run *Run) GetOrCreateLocation(location *Location) (int, error) {
	for i, l := range run.Artifacts {
		if l.Location.Uri == location.Uri {
			return i, nil
		}
	}
	run.Artifacts = append(run.Artifacts, &LocationWrapper{
		Location: &Location{
			Uri: location.Uri,
		},
	})
	return len(run.Artifacts) - 1, nil
}

// AddResultDetails adds rules to the driver and artifact locations if they are missing. It adds the result to the result block as well
func (run *Run) AddResultDetails(rule *Rule, location *PhysicalLocation, result *Result) error {
	ruleIndex, err := run.Tool.Driver.GetOrCreateRule(rule)
	if err != nil {
		return err
	}
	result.RuleIndex = ruleIndex
	locationIndex, err := run.GetOrCreateLocation(&Location{Uri: location.ArtifactLocation.Uri})
	if err != nil {
		return nil
	}
	location.ArtifactLocation.Index = locationIndex
	result.AddLocation(location)

	run.Results = append(run.Results, result)
	return nil
}
