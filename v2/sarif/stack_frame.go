package sarif

// StackFrame ...
type StackFrame struct {
	PropertyBag
	Location   *Location `json:"location,omitempty"`
	Module     *string   `json:"module,omitempty"`
	Parameters []string  `json:"parameters,omitempty"`
	ThreadID   *int      `json:"threadId,omitempty"`
}

// NewStackFrame ...
func NewStackFrame() *StackFrame {
	return &StackFrame{}
}

// WithLocation sets the Location
func (stackFrame *StackFrame) WithLocation(location *Location) *StackFrame {
	stackFrame.Location = location
	return stackFrame
}

// WithModule sets the Module
func (stackFrame *StackFrame) WithModule(module string) *StackFrame {
	stackFrame.Module = &module
	return stackFrame
}

// WithParameters sets the Parameters
func (stackFrame *StackFrame) WithParameters(parameters []string) *StackFrame {
	stackFrame.Parameters = parameters
	return stackFrame
}

// AddParameter ...
func (stackFrame *StackFrame) AddParameter(parameter string) {
	stackFrame.Parameters = append(stackFrame.Parameters, parameter)
}

// WithThreadID sets the ThreadID
func (stackFrame *StackFrame) WithThreadID(threadID int) *StackFrame {
	stackFrame.ThreadID = &threadID
	return stackFrame
}
