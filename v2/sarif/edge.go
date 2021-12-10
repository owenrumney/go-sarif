package sarif

type Edge struct {
	PropertyBag
	ID           string   `json:"id"`
	Label        *Message `json:"label,omitempty"`
	SourceNodeID string   `json:"sourceNodeId"`
	TargetNodeID string   `json:"targetNodeId"`
}

func NewEdge(id, sourceNodeID, targetNodeID string) *Edge {
	return &Edge{
		ID:           id,
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}

func (edge *Edge) WithID(id string) *Edge {
	edge.ID = id
	return edge
}

func (edge *Edge) WithLabel(label *Message) *Edge {
	edge.Label = label
	return edge
}

func (edge *Edge) WithLabelText(text string) *Edge {
	edge.Label = &Message{
		Text: &text,
	}
	return edge
}

func (edge *Edge) WithLabelMarkdown(markdown string) *Edge {
	edge.Label = &Message{
		Markdown: &markdown,
	}
	return edge
}
