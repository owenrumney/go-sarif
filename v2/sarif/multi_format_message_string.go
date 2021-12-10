package sarif

type MultiformatMessageString struct {
	PropertyBag
	Text     *string `json:"text,omitempty"`
	Markdown *string `json:"markdown,omitempty"`
}

func NewMarkdownMultiformatMessageString(markdown string) *MultiformatMessageString {
	return &MultiformatMessageString{
		Markdown: &markdown,
	}
}

func NewMultiformatMessageString(text string) *MultiformatMessageString {
	return &MultiformatMessageString{
		Text: &text,
	}
}

func (multiFormatMessageString *MultiformatMessageString) WithText(text string) *MultiformatMessageString {
	multiFormatMessageString.Text = &text
	return multiFormatMessageString
}
func (multiFormatMessageString *MultiformatMessageString) WithMarkdown(markdown string) *MultiformatMessageString {
	multiFormatMessageString.Markdown = &markdown
	return multiFormatMessageString
}
