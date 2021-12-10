package sarif

type Graph struct {
	PropertyBag
	Description *Message `json:"description,omitempty"`
	Edges       []*Edge  `json:"edges,omitempty"`
	Nodes       []*Node  `json:"nodes,omitempty"`
}

func NewGraph() *Graph {
	return &Graph{}
}

func (graph *Graph) WithDescription(message *Message) *Graph {
	graph.Description = message
	return graph
}

func (graph *Graph) WithDescriptionText(text string) *Graph {
	if graph.Description == nil {
		graph.Description = &Message{}
	}
	graph.Description.Text = &text
	return graph
}

func (graph *Graph) WithDescriptionMarkdown(markdown string) *Graph {
	if graph.Description == nil {
		graph.Description = &Message{}
	}
	graph.Description.Markdown = &markdown
	return graph
}

func (graph *Graph) WithEdges(edges []*Edge) *Graph {
	graph.Edges = edges
	return graph
}

func (graph *Graph) AddEdge(edge *Edge) {
	graph.Edges = append(graph.Edges, edge)
}

func (graph *Graph) WithNodes(nodes []*Node) *Graph {
	graph.Nodes = nodes
	return graph
}

func (graph *Graph) AddNode(node *Node) {
	graph.Nodes = append(graph.Nodes, node)
}
