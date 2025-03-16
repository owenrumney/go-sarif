package sarif

// AddRule returns an existing ReportingDescriptor for the ruleID or creates a new ReportingDescriptor and returns a pointer to it
func (run *Run) AddRule(ruleID string) *ReportingDescriptor {
	for _, rule := range run.Tool.Driver.Rules {
		if rule.ID == ruleID {
			return rule
		}
	}
	rule := NewRule(ruleID)
	run.Tool.Driver.Rules = append(run.Tool.Driver.Rules, rule)
	return rule
}

// AddDistinctArtifact will handle deduplication of simple artifact additions
func (run *Run) AddDistinctArtifact(uri string) *Artifact {
	for _, artifact := range run.Artifacts {
		if artifact.Location.URI == uri {
			return artifact
		}
	}

	a := NewArtifact().WithLength(-1)
	a.WithLocation(NewSimpleArtifactLocation(uri))

	run.Artifacts = append(run.Artifacts, a)
	return a
}

// CreateResultForRule returns an existing Result or creates a new one and returns a pointer to it
func (run *Run) CreateResultForRule(ruleID string) *Result {
	result := NewRuleResult(ruleID)
	run.AddResult(result)
	return result
}

// DedupeArtifacts will remove any duplicate artifacts from the run
func (run *Run) DedupeArtifacts() error {
	dupes := map[*Artifact]bool{}
	deduped := []*Artifact{}

	for _, a := range run.Artifacts {
		if _, ok := dupes[a]; !ok {
			dupes[a] = true
			deduped = append(deduped, a)
		}
	}
	run.Artifacts = deduped
	return nil
}
