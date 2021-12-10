package sarif

type ThreadFlow struct {
	PropertyBag
	ID             *string                              `json:"id,omitempty"`
	ImmutableState map[string]*MultiformatMessageString `json:"immutableState,omitempty"`
	InitialState   map[string]*MultiformatMessageString `json:"initialState,omitempty"`
	Locations      []*ThreadFlowLocation                `json:"locations"`
	Message        *Message                             `json:"message,omitempty"`
}

func NewThreadFlow() *ThreadFlow {
	return &ThreadFlow{}
}

func (threadFlow *ThreadFlow) WithID(id string) *ThreadFlow {
	threadFlow.ID = &id
	return threadFlow
}

func (threadFlow *ThreadFlow) WithImmutableState(immutableState map[string]*MultiformatMessageString) *ThreadFlow {
	threadFlow.ImmutableState = immutableState
	return threadFlow
}

func (threadFlow *ThreadFlow) WithInitialState(initialState map[string]*MultiformatMessageString) *ThreadFlow {
	threadFlow.InitialState = initialState
	return threadFlow
}

func (threadFlow *ThreadFlow) WithLocations(locations []*ThreadFlowLocation) *ThreadFlow {
	threadFlow.Locations = locations
	return threadFlow
}

func (threadFlow *ThreadFlow) AddLocation(location *ThreadFlowLocation) {
	threadFlow.Locations = append(threadFlow.Locations, location)
}


func (threadFlow *ThreadFlow) WithMessage(message *Message) *ThreadFlow {
	threadFlow.Message = message
	return threadFlow
}

func (threadFlow *ThreadFlow) WithTextMessage(text string) *ThreadFlow {
	if threadFlow.Message == nil {
		threadFlow.Message = &Message{}
	}
	threadFlow.Message.Text = &text
	return threadFlow
}

func (threadFlow *ThreadFlow) WithMessageMarkdown(markdown string) *ThreadFlow {
	if threadFlow.Message == nil {
		threadFlow.Message = &Message{}
	}
	threadFlow.Message.Markdown = &markdown
	return threadFlow
}
