package sarif

type Attachment struct {
	PropertyBag
	ArtifactLocation *ArtifactLocation `json:"artifactLocation,omitempty"`
	Description      *Message          `json:"description,omitempty"`
	Rectangles       []*Rectangle      `json:"rectangles,omitempty"`
}

func NewAttachment() *Attachment {
	return &Attachment{}
}

func (attachment *Attachment) WithArtifactionLocation(artifactLocation *ArtifactLocation) *Attachment {
	attachment.ArtifactLocation = artifactLocation
	return attachment
}

func (attachment *Attachment) WithDescription(description *Message) *Attachment {
	attachment.Description = description
	return attachment
}

func (attachment *Attachment) WithDescriptionText(text string) *Attachment {
	if attachment.Description == nil {
		attachment.Description = &Message{}
	}
	attachment.Description.Text = &text
	return attachment
}

func (attachment *Attachment) WithDescriptionMarkdown(markdown string) *Attachment {
	if attachment.Description == nil {
		attachment.Description = &Message{}
	}
	attachment.Description.Markdown = &markdown
	return attachment
}

func (attachment *Attachment) WithRectangles(rectangles []*Rectangle) *Attachment {
	attachment.Rectangles = rectangles
	return attachment
}

func (attachment *Attachment) AddRectangle(rectangle *Rectangle) {
	attachment.Rectangles = append(attachment.Rectangles, rectangle)
}
