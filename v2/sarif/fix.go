package sarif

type Fix struct {
	PropertyBag
	Description     *Message          `json:"description,omitempty"`
	ArtifactChanges []*ArtifactChange `json:"artifactChanges"` //	required
}

func NewFix() *Fix {
	return &Fix{}
}

func (fix *Fix) WithDescription(message *Message) *Fix {
	fix.Description = message
	return fix
}

func (fix *Fix) WithDescriptionText(text string) *Fix {
	if fix.Description == nil {
		fix.Description = &Message{}
	}
	fix.Description.Text = &text
	return fix
}

func (fix *Fix) WithDescriptionMarkdown(markdown string) *Fix {
	if fix.Description == nil {
		fix.Description = &Message{}
	}
	fix.Description.Markdown = &markdown
	return fix
}

func (fix *Fix) WithArtifactChanges(artifactChanges []*ArtifactChange) *Fix {
	fix.ArtifactChanges = artifactChanges
	return fix
}

func (fix *Fix) AddArtifactChanges(artifactChange *ArtifactChange) {
	fix.ArtifactChanges = append(fix.ArtifactChanges, artifactChange)
}
