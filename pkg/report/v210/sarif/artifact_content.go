package sarif

// ArtifactContent - Represents the contents of an artifact.
type ArtifactContent struct {
	// UTF-8-encoded content from a text artifact.
	Text string `json:"text,omitempty"`

	// MIME Base64-encoded content from a binary artifact, or from a text artifact in its original encoding.
	Binary string `json:"binary,omitempty"`

	// An alternate rendered representation of the artifact (e.g., a decompiled representation of a binary region).
	Rendered *MultiformatMessageString `json:"rendered,omitempty"`

	// Key/value pairs that provide additional information about the artifact content.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// NewArtifactContent - creates a new
func NewArtifactContent() *ArtifactContent {
	return &ArtifactContent{}
}

// WithText - add a Text to the ArtifactContent
func (t *ArtifactContent) WithText(text string) *ArtifactContent {
	t.Text = text
	return t
}

// WithBinary - add a Binary to the ArtifactContent
func (b *ArtifactContent) WithBinary(binary string) *ArtifactContent {
	b.Binary = binary
	return b
}

// WithRendered - add a Rendered to the ArtifactContent
func (r *ArtifactContent) WithRendered(rendered *MultiformatMessageString) *ArtifactContent {
	r.Rendered = rendered
	return r
}

// WithProperties - add a Properties to the ArtifactContent
func (p *ArtifactContent) WithProperties(properties *PropertyBag) *ArtifactContent {
	p.Properties = properties
	return p
}
