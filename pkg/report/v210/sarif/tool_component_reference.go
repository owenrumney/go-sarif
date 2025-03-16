package sarif

// ToolComponentReference - Identifies a particular toolComponent object, either the driver or an extension.
type ToolComponentReference struct {
	// The 'name' property of the referenced toolComponent.
	Name string `json:"name,omitempty"`

	// An index into the referenced toolComponent in tool.extensions.
	Index int `json:"index,omitempty"`

	// The 'guid' property of the referenced toolComponent.
	GuID string `json:"guid,omitempty"`

	// Key/value pairs that provide additional information about the toolComponentReference.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// NewToolComponentReference - creates a new
func NewToolComponentReference() *ToolComponentReference {
	return &ToolComponentReference{}
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

// WithGuID - add a GuID to the ToolComponentReference
func (g *ToolComponentReference) WithGuID(guid string) *ToolComponentReference {
	g.GuID = guid
	return g
}

// WithProperties - add a Properties to the ToolComponentReference
func (p *ToolComponentReference) WithProperties(properties *PropertyBag) *ToolComponentReference {
	p.Properties = properties
	return p
}
