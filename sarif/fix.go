package sarif

type Fix struct {
	Description     *Message          `json:"description,omitempty"`
	ArtifactChanges []*ArtifactChange `json:"artifactChanges"` //	required
}
