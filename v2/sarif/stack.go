package sarif

// Stack ...
type Stack struct {
	PropertyBag
	Frames     []*StackFrame `json:"frames"`
	Message    *Message      `json:"message,omitempty"`
	Properties *PropertyBag  `json:"properties,omitempty"`
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{}
}

// WithFrames sets the Frames
func (stack *Stack) WithFrames(frames []*StackFrame) *Stack {
	stack.Frames = frames
	return stack
}

// AddFrame ...
func (stack *Stack) AddFrame(frame *StackFrame) {
	stack.Frames = append(stack.Frames, frame)
}

// WithMessage sets the Message
func (stack *Stack) WithMessage(message *Message) *Stack {
	stack.Message = message
	return stack
}

// WithTextMessage sets the TextMessage
func (stack *Stack) WithTextMessage(text string) *Stack {
	if stack.Message == nil {
		stack.Message = &Message{}
	}
	stack.Message.Text = &text
	return stack
}

// WithMessageMarkdown sets the MessageMarkdown
func (stack *Stack) WithMessageMarkdown(markdown string) *Stack {
	if stack.Message == nil {
		stack.Message = &Message{}
	}
	stack.Message.Markdown = &markdown
	return stack
}
