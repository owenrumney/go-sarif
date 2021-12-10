package sarif

type EdgeTraversal struct {
	PropertyBag
	EdgeID            string                               `json:"edgeId"`
	FinalState        map[string]*MultiformatMessageString `json:"finalState,omitempty"`
	Message           *Message                             `json:"message,omitempty"`
	StepOverEdgeCount *int                                 `json:"stepOverEdgeCount,omitempty"`
}

func NewEdgeTraversal(edgeID string) *EdgeTraversal {
	return &EdgeTraversal{
		EdgeID: edgeID,
	}
}

func (edgeTraversal *EdgeTraversal) WithDescription(message *Message) *EdgeTraversal {
	edgeTraversal.Message = message
	return edgeTraversal
}

func (edgeTraversal *EdgeTraversal) WithDescriptionText(text string) *EdgeTraversal {
	edgeTraversal.Message = &Message{
		Text: &text,
	}
	return edgeTraversal
}

func (edgeTraversal *EdgeTraversal) WithDescriptionMarkdown(markdown string) *EdgeTraversal {
	edgeTraversal.Message = &Message{
		Markdown: &markdown,
	}
	return edgeTraversal
}

func (edgeTraversal *EdgeTraversal) WithFinalState(finalState map[string]*MultiformatMessageString) *EdgeTraversal {
	edgeTraversal.FinalState = finalState
	return edgeTraversal
}

func (edgeTraversal *EdgeTraversal) SetFinalState(key string, state *MultiformatMessageString) {
	edgeTraversal.FinalState[key] = state
}

func (edgeTraversal *EdgeTraversal) WithStepOverEdgeCount(stepOverEdgeCount int) *EdgeTraversal {
	edgeTraversal.StepOverEdgeCount = &stepOverEdgeCount
	return edgeTraversal
}