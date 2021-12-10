package sarif

type Message struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540897
	PropertyBag
	Text      *string  `json:"text,omitempty"`
	Markdown  *string  `json:"markdown,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}

func NewMessage() *Message {
	return &Message{}
}

func NewTextMessage(text string) *Message {
	return NewMessage().WithText(text)
}

func NewMarkdownMessage(markdown string) *Message {
	return NewMessage().WithMarkdown(markdown)
}

func (message *Message) WithText(text string) *Message {
	message.Text = &text
	return message
}

func (message *Message) WithMarkdown(markdown string) *Message {
	message.Markdown = &markdown
	return message
}

func (message *Message) WithId(id string) *Message {
	message.Id = &id
	return message
}

func (message *Message) WithArguments(arguments []string) *Message {
	message.Arguments = arguments
	return message
}

func (message *Message) AddArgument(argument string) {
	message.Arguments = append(message.Arguments, argument)
}
