package sarif

type Node struct {
	PropertyBag
	Children []*Node   `json:"children,omitempty"`
	ID       string    `json:"id"`
	Label    *Message  `json:"label,omitempty"`
	Location *Location `json:"location,omitempty"`
}

func NewNode(id string) *Node {
	return &Node{
		ID: id,
	}
}

func (node *Node) WithChildren(children []*Node) *Node {
	node.Children = children
	return node
}

func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
}

func (node *Node) WithLabel(message *Message) *Node {
	node.Label = message
	return node
}

func (node *Node) WithLabelText(text string) *Node {
	if node.Label == nil {
		node.Label = &Message{}
	}
	node.Label.Text = &text
	return node
}

func (node *Node) WithLabelMarkdown(markdown string) *Node {
	if node.Label == nil {
		node.Label = &Message{}
	}
	node.Label.Markdown = &markdown
	return node
}
