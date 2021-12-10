package sarif

// ArtifactChange ...
type ArtifactChange struct {
	PropertyBag
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Replacements     []*Replacement   `json:"replacements"`
}

// NewArtifactChange ...
func NewArtifactChange(artifactLocation *ArtifactLocation) *ArtifactChange {
	return &ArtifactChange{
		ArtifactLocation: *artifactLocation,
	}
}

// WithReplacement sets the Replacement
func (artifactChange *ArtifactChange) WithReplacement(replacement *Replacement) *ArtifactChange {
	artifactChange.Replacements = append(artifactChange.Replacements, replacement)
	return artifactChange
}
