package sarif

// MultiformatMessageString ...
type MultiformatMessageString struct {
	PropertyBag
	Text     *string `json:"text,omitempty"`
	Markdown *string `json:"markdown,omitempty"`
}

// NewMarkdownMultiformatMessageString ...
func NewMarkdownMultiformatMessageString(markdown string) *MultiformatMessageString {
	return &MultiformatMessageString{
		Markdown: &markdown,
	}
}

// NewMultiformatMessageString ...
func NewMultiformatMessageString(text string) *MultiformatMessageString {
	return &MultiformatMessageString{
		Text: &text,
	}
}

// WithText sets the Text
func (multiFormatMessageString *MultiformatMessageString) WithText(text string) *MultiformatMessageString {
	multiFormatMessageString.Text = &text
	return multiFormatMessageString
}
// WithMarkdown sets the Markdown
func (multiFormatMessageString *MultiformatMessageString) WithMarkdown(markdown string) *MultiformatMessageString {
	multiFormatMessageString.Markdown = &markdown
	return multiFormatMessageString
}
