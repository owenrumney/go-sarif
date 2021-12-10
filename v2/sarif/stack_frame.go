package sarif

type StackFrame struct {
	PropertyBag
	Location   *Location `json:"location,omitempty"`
	Module     *string   `json:"module,omitempty"`
	Parameters []string  `json:"parameters,omitempty"`
	ThreadID   *int      `json:"threadId,omitempty"`
}

func NewStackFrame() *StackFrame {
	return &StackFrame{}
}

func (stackFrame *StackFrame) WithLocation(location *Location) *StackFrame {
	stackFrame.Location = location
	return stackFrame
}

func (stackFrame *StackFrame) WithModule(module string) *StackFrame {
	stackFrame.Module = &module
	return stackFrame
}

func (stackFrame *StackFrame) WithParameters(parameters []string) *StackFrame {
	stackFrame.Parameters = parameters
	return stackFrame
}

func (stackFrame *StackFrame) AddParameter(parameter string) {
	stackFrame.Parameters = append(stackFrame.Parameters, parameter)
}

func (stackFrame *StackFrame) WithThreadID(threadID int) *StackFrame {
	stackFrame.ThreadID = &threadID
	return stackFrame
}
