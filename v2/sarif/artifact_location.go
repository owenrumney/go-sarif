package sarif

type ArtifactLocation struct {
	PropertyBag
	URI         *string  `json:"uri,omitempty"`
	URIBaseId   *string  `json:"uriBaseId,omitempty"`
	Index       *uint    `json:"index,omitempty"`
	Description *Message `json:"description,omitempty"`
}

func NewArtifactLocation() *ArtifactLocation {
	return &ArtifactLocation{}
}

func NewSimpleArtifactLocation(uri string) *ArtifactLocation {
	return NewArtifactLocation().WithUri(uri)
}

func (artifactLocation *ArtifactLocation) WithUri(uri string) *ArtifactLocation {
	artifactLocation.URI = &uri
	return artifactLocation
}

func (artifactLocation *ArtifactLocation) WithUriBaseId(uriBaseId string) *ArtifactLocation {
	artifactLocation.URIBaseId = &uriBaseId
	return artifactLocation
}

func (artifactLocation *ArtifactLocation) WithIndex(index int) *ArtifactLocation {
	i := uint(index)
	artifactLocation.Index = &i
	return artifactLocation
}

func (artifactLocation *ArtifactLocation) WithDescription(message *Message) *ArtifactLocation {
	artifactLocation.Description = message
	return artifactLocation
}

func (artifactLocation *ArtifactLocation) WithDescriptionText(text string) *ArtifactLocation {
	if artifactLocation.Description == nil {
		artifactLocation.Description = &Message{}
	}
	artifactLocation.Description.Text = &text
	return artifactLocation
}

func (artifactLocation *ArtifactLocation) WithDescriptionMarkdown(markdown string) *ArtifactLocation {
	if artifactLocation.Description == nil {
		artifactLocation.Description = &Message{}
	}
	artifactLocation.Description.Markdown = &markdown
	return artifactLocation
}
