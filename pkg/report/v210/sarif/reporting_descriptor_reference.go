package sarif

// ReportingDescriptorReference - Information about how to locate a relevant reporting descriptor.
type ReportingDescriptorReference struct {
	// The id of the descriptor.
	ID string `json:"id,omitempty"`

	// The index into an array of descriptors in toolComponent.ruleDescriptors, toolComponent.notificationDescriptors, or toolComponent.taxonomyDescriptors, depending on context.
	Index int `json:"index,omitempty"`

	// A guid that uniquely identifies the descriptor.
	GuID string `json:"guid,omitempty"`

	// A reference used to locate the toolComponent associated with the descriptor.
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`

	// Key/value pairs that provide additional information about the reporting descriptor reference.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// NewReportingDescriptorReference - creates a new
func NewReportingDescriptorReference() *ReportingDescriptorReference {
	return &ReportingDescriptorReference{}
}

// WithID - add a ID to the ReportingDescriptorReference
func (i *ReportingDescriptorReference) WithID(id string) *ReportingDescriptorReference {
	i.ID = id
	return i
}

// WithIndex - add a Index to the ReportingDescriptorReference
func (i *ReportingDescriptorReference) WithIndex(index int) *ReportingDescriptorReference {
	i.Index = index
	return i
}

// WithGuID - add a GuID to the ReportingDescriptorReference
func (g *ReportingDescriptorReference) WithGuID(guid string) *ReportingDescriptorReference {
	g.GuID = guid
	return g
}

// WithToolComponent - add a ToolComponent to the ReportingDescriptorReference
func (t *ReportingDescriptorReference) WithToolComponent(toolComponent *ToolComponentReference) *ReportingDescriptorReference {
	t.ToolComponent = toolComponent
	return t
}

// WithProperties - add a Properties to the ReportingDescriptorReference
func (p *ReportingDescriptorReference) WithProperties(properties *PropertyBag) *ReportingDescriptorReference {
	p.Properties = properties
	return p
}
