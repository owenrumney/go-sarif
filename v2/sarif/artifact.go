package sarif

type Artifact struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10541049
	PropertyBag
	Location            *ArtifactLocation `json:"location,omitempty"`
	ParentIndex         *uint             `json:"parentIndex,omitempty"`
	Offset              *uint             `json:"offset,omitempty"`
	Length              int               `json:"length"`
	Roles               []string          `json:"roles,omitempty"`
	MimeType            *string           `json:"mimeType,omitempty"`
	Contents            *ArtifactContent  `json:"contents,omitempty"`
	Encoding            *string           `json:"encoding,omitempty"`
	SourceLanguage      *string           `json:"sourceLanguage,omitempty"`
	Hashes              map[string]string `json:"hashes,omitempty"`
	LastModifiedTimeUtc *string           `json:"lastModifiedTimeUtc,omitempty"`
	Description         *Message          `json:"description,omitempty"`
}

func NewArtifact() *Artifact {
	return &Artifact{}
}

func (artifact *Artifact) WithLocation(artifactLocation *ArtifactLocation) *Artifact {
	artifact.Location = artifactLocation
	return artifact
}

func (artifact *Artifact) WithParentIndex(parentIndex int) *Artifact {
	i := uint(parentIndex)
	artifact.ParentIndex = &i
	return artifact
}

func (artifact *Artifact) WithOffset(offset int) *Artifact {
	o := uint(offset)
	artifact.Offset = &o
	return artifact
}

func (artifact *Artifact) WithLength(length int) *Artifact {
	artifact.Length = length
	return artifact
}

func (artifact *Artifact) WithRole(role string) *Artifact {
	artifact.Roles = append(artifact.Roles, role)
	return artifact
}

func (artifact *Artifact) WithMimeType(mimeType string) *Artifact {
	artifact.MimeType = &mimeType
	return artifact
}

func (artifact *Artifact) WithContents(artifactContent *ArtifactContent) *Artifact {
	artifact.Contents = artifactContent
	return artifact
}

func (artifact *Artifact) WithEncoding(encoding string) *Artifact {
	artifact.Encoding = &encoding
	return artifact
}

func (artifact *Artifact) WithSourceLanguage(sourceLanguage string) *Artifact {
	artifact.SourceLanguage = &sourceLanguage
	return artifact
}

func (artifact *Artifact) WithHashes(hashes map[string]string) *Artifact {
	artifact.Hashes = hashes
	return artifact
}

func (artifact *Artifact) WithLastModifiedTimeUtc(lastModified string) *Artifact {
	artifact.LastModifiedTimeUtc = &lastModified
	return artifact
}

func (artifact *Artifact) WithDescription(message *Message) *Artifact {
	artifact.Description = message
	return artifact
}

func (artifact *Artifact) WithDescriptionText(text string) *Artifact {
	if artifact.Description == nil {
		artifact.Description = &Message{}
	}
	artifact.Description.Text = &text
	return artifact
}

func (artifact *Artifact) WithDescriptionMarkdown(markdown string) *Artifact {
	if artifact.Description == nil {
		artifact.Description = &Message{}
	}
	artifact.Description.Markdown = &markdown
	return artifact
}
