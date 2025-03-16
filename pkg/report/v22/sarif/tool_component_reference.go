package sarif

// ToolComponentReference - Identifies a particular toolComponent object, either the driver or an extension.
type ToolComponentReference struct {
	// Key/value pairs that provide additional information about the toolComponentReference.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The 'name' property of the referenced toolComponent.
	Name string `json:"name,omitempty"`

	// An index into the referenced toolComponent in tool.extensions.
	Index int `json:"index,omitempty"`

	// The 'guid' property of the referenced toolComponent.
	Guid *Guid `json:"guid,omitempty"`
}

// NewToolComponentReference - creates a new
func NewToolComponentReference() *ToolComponentReference {
	return &ToolComponentReference{}
}

// WithProperties - add a Properties to the ToolComponentReference
func (p *ToolComponentReference) WithProperties(properties *PropertyBag) *ToolComponentReference {
	p.Properties = properties
	return p
}

// WithName - add a Name to the ToolComponentReference
func (n *ToolComponentReference) WithName(name string) *ToolComponentReference {
	n.Name = name
	return n
}

// WithIndex - add a Index to the ToolComponentReference
func (i *ToolComponentReference) WithIndex(index int) *ToolComponentReference {
	i.Index = index
	return i
}

// WithGuid - add a Guid to the ToolComponentReference
func (g *ToolComponentReference) WithGuid(guid *Guid) *ToolComponentReference {
	g.Guid = guid
	return g
}
