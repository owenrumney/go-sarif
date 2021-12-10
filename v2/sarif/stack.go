package sarif

type Stack struct {
	PropertyBag
	Frames     []*StackFrame `json:"frames"`
	Message    *Message      `json:"message,omitempty"`
	Properties *PropertyBag  `json:"properties,omitempty"`
}


func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) WithFrames (frames []*StackFrame) *Stack {
	stack.Frames = frames
	return stack
}

func (stack *Stack) AddFrame(frame *StackFrame) {
	stack.Frames = append(stack.Frames, frame)
}

func (stack *Stack) WithMessage(message *Message) *Stack {
	stack.Message = message
	return stack
}

func (stack *Stack) WithTextMessage(text string) *Stack {
	if stack.Message == nil {
		stack.Message = &Message{}
	}
	stack.Message.Text = &text
	return stack
}

func (stack *Stack) WithMessageMarkdown(markdown string) *Stack {
	if stack.Message == nil {
		stack.Message = &Message{}
	}
	stack.Message.Markdown = &markdown
	return stack
}