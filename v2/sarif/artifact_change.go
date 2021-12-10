package sarif

type ArtifactChange struct {
	PropertyBag
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Replacements     []*Replacement   `json:"replacements"`
}

func NewArtifactChange(artifactLocation *ArtifactLocation) *ArtifactChange {
	return &ArtifactChange{
		ArtifactLocation: *artifactLocation,
	}
}

func (artifactChange *ArtifactChange) WithReplacement(replacement *Replacement) *ArtifactChange {
	artifactChange.Replacements = append(artifactChange.Replacements, replacement)
	return artifactChange
}
