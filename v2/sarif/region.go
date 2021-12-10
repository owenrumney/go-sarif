package sarif

type Region struct {
	PropertyBag
	StartLine      *int             `json:"startLine,omitempty"`
	StartColumn    *int             `json:"startColumn,omitempty"`
	EndLine        *int             `json:"endLine,omitempty"`
	EndColumn      *int             `json:"endColumn,omitempty"`
	CharOffset     *int             `json:"charOffset,omitempty"`
	CharLength     *int             `json:"charLength,omitempty"`
	ByteOffset     *int             `json:"byteOffset,omitempty"`
	ByteLength     *int             `json:"byteLength,omitempty"`
	Snippet        *ArtifactContent `json:"snippet,omitempty"`
	Message        *Message         `json:"message,omitempty"`
	SourceLanguage *string          `json:"sourceLanguage,omitempty"`
}

func NewRegion() *Region {
	return &Region{}
}

func NewSimpleRegion(startLine, endLine int) *Region {
	return NewRegion().
		WithStartLine(startLine).
		WithEndLine(endLine)
}

func (region *Region) WithStartLine(startLine int) *Region {
	region.StartLine = &startLine
	return region
}

func (region *Region) WithStartColumn(startColumn int) *Region {
	region.StartColumn = &startColumn
	return region
}

func (region *Region) WithEndLine(endLine int) *Region {
	region.EndLine = &endLine
	return region
}

func (region *Region) WithEndColumn(endColumn int) *Region {
	region.EndColumn = &endColumn
	return region
}

func (region *Region) WithCharOffset(charOffset int) *Region {
	region.CharOffset = &charOffset
	return region
}

func (region *Region) WithCharLength(charLength int) *Region {
	region.CharLength = &charLength
	return region
}

func (region *Region) WithByteOffset(byteOffset int) *Region {
	region.ByteOffset = &byteOffset
	return region
}

func (region *Region) WithByteLength(byteLength int) *Region {
	region.ByteLength = &byteLength
	return region
}

func (region *Region) WithSnippet(snippet *ArtifactContent) *Region {
	region.Snippet = snippet
	return region
}

func (region *Region) WithMessage(message *Message) *Region {
	region.Message = message
	return region
}

func (region *Region) WithTextMessage(text string) *Region {
	if region.Message == nil {
		region.Message = &Message{}
	}
	region.Message.Text = &text
	return region
}

func (region *Region) WithMessageMarkdown(markdown string) *Region {
	if region.Message == nil {
		region.Message = &Message{}
	}
	region.Message.Markdown = &markdown
	return region
}

func (region *Region) WithSourceLanguage(sourceLanguage string) *Region {

	return region
}
