package sarif

type CodeFlow struct {
	PropertyBag
	Message     *Message      `json:"message,omitempty"`
	ThreadFlows []*ThreadFlow `json:"threadFlows,omitempty"`
}

func NewCodeFlow() *CodeFlow {
	return &CodeFlow{}
}

func (codeFlow *CodeFlow) WithMessage(message *Message) *CodeFlow {
	codeFlow.Message = message
	return codeFlow
}

func (codeFlow *CodeFlow) WithTextMessage(text string) *CodeFlow {
	if codeFlow.Message == nil {
		codeFlow.Message = &Message{}
	}
	codeFlow.Message.Text = &text
	return codeFlow
}

func (codeFlow *CodeFlow) WithMessageMarkdown(markdown string) *CodeFlow {
	if codeFlow.Message == nil {
		codeFlow.Message = &Message{}
	}
	codeFlow.Message.Markdown = &markdown
	return codeFlow
}

func (codeFlow *CodeFlow) WithThreadFlows(threadFlows []*ThreadFlow) *CodeFlow {
	codeFlow.ThreadFlows = threadFlows

	return codeFlow
}

func (codeFlow *CodeFlow) AddThreadFlow(threadFlow *ThreadFlow) {
	codeFlow.ThreadFlows = append(codeFlow.ThreadFlows, threadFlow)
}
