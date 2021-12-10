package sarif

// Message ...
type Message struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540897
	PropertyBag
	Text      *string  `json:"text,omitempty"`
	Markdown  *string  `json:"markdown,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}

// NewMessage ...
func NewMessage() *Message {
	return &Message{}
}

// NewTextMessage ...
func NewTextMessage(text string) *Message {
	return NewMessage().WithText(text)
}

// NewMarkdownMessage ...
func NewMarkdownMessage(markdown string) *Message {
	return NewMessage().WithMarkdown(markdown)
}

// WithText sets the Text
func (message *Message) WithText(text string) *Message {
	message.Text = &text
	return message
}

// WithMarkdown sets the Markdown
func (message *Message) WithMarkdown(markdown string) *Message {
	message.Markdown = &markdown
	return message
}

// WithId sets the Id
func (message *Message) WithId(id string) *Message {
	message.Id = &id
	return message
}

// WithArguments sets the Arguments
func (message *Message) WithArguments(arguments []string) *Message {
	message.Arguments = arguments
	return message
}

// AddArgument ...
func (message *Message) AddArgument(argument string) {
	message.Arguments = append(message.Arguments, argument)
}
