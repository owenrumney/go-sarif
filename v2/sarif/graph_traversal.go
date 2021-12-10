package sarif

type GraphTraversal struct {
	PropertyBag
	Description      *Message                             `json:"description,omitempty"`
	EdgeTraversals   []*EdgeTraversal                     `json:"edgeTraversals,omitempty"`
	ImmutableState   map[string]*MultiformatMessageString `json:"immutableState,omitempty"`
	InitialState     map[string]*MultiformatMessageString `json:"initialState,omitempty"`
	ResultGraphIndex *int                                 `json:"resultGraphIndex,omitempty"`
	RunGraphIndex    *int                                 `json:"runGraphIndex,omitempty"`
}

func NewGraphTraversal() *GraphTraversal {
	return &GraphTraversal{}
}

func (graphTraversal *GraphTraversal) WithDescription(message *Message) *GraphTraversal {
	graphTraversal.Description = message
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithDescriptionText(text string) *GraphTraversal {
	if graphTraversal.Description == nil {
		graphTraversal.Description = &Message{}
	}
	graphTraversal.Description.Text = &text
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithDescriptionMarkdown(markdown string) *GraphTraversal {
	if graphTraversal.Description == nil {
		graphTraversal.Description = &Message{}
	}
	graphTraversal.Description.Markdown = &markdown
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithEdgeTraversals(edgeTraversals []*EdgeTraversal) *GraphTraversal {
	graphTraversal.EdgeTraversals = edgeTraversals
	return graphTraversal
}

func (graphTraversal *GraphTraversal) AddEdge(edgeTraversal *EdgeTraversal) {
	graphTraversal.EdgeTraversals = append(graphTraversal.EdgeTraversals, edgeTraversal)
}

func (graphTraversal *GraphTraversal) WithResultGraphIndex(index int) *GraphTraversal {
	graphTraversal.ResultGraphIndex = &index
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithRunGraphIndex(index int) *GraphTraversal {
	graphTraversal.RunGraphIndex = &index
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithImmutableState(immutableState map[string]*MultiformatMessageString) *GraphTraversal {
	graphTraversal.ImmutableState = immutableState
	return graphTraversal
}

func (graphTraversal *GraphTraversal) WithInitialState(initialState map[string]*MultiformatMessageString) *GraphTraversal {
	graphTraversal.InitialState = initialState
	return graphTraversal
}