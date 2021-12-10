package sarif

type ArtifactContent struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540860
	PropertyBag
	Text     *string                   `json:"text,omitempty"`
	Binary   *string                   `json:"binary,omitempty"`
	Rendered *MultiformatMessageString `json:"rendered,omitempty"`
}

func NewArtifactContent() *ArtifactContent {
	return &ArtifactContent{}
}

func (artifactContent *ArtifactContent) WithText(text string) *ArtifactContent {
	artifactContent.Text = &text
	return artifactContent
}

func (artifactContent *ArtifactContent) WithBinary(binary string) *ArtifactContent {
	artifactContent.Binary = &binary
	return artifactContent
}

func (artifactContent *ArtifactContent) WithRendered(mms *MultiformatMessageString) *ArtifactContent {
	artifactContent.Rendered = mms
	return artifactContent
}
