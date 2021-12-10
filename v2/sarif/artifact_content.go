package sarif

// ArtifactContent ...
type ArtifactContent struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540860
	PropertyBag
	Text     *string                   `json:"text,omitempty"`
	Binary   *string                   `json:"binary,omitempty"`
	Rendered *MultiformatMessageString `json:"rendered,omitempty"`
}

// NewArtifactContent ...
func NewArtifactContent() *ArtifactContent {
	return &ArtifactContent{}
}

// WithText sets the Text
func (artifactContent *ArtifactContent) WithText(text string) *ArtifactContent {
	artifactContent.Text = &text
	return artifactContent
}

// WithBinary sets the Binary
func (artifactContent *ArtifactContent) WithBinary(binary string) *ArtifactContent {
	artifactContent.Binary = &binary
	return artifactContent
}

// WithRendered sets the Rendered
func (artifactContent *ArtifactContent) WithRendered(mms *MultiformatMessageString) *ArtifactContent {
	artifactContent.Rendered = mms
	return artifactContent
}
