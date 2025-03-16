package sarif

// RunAutomationDetails - Information that describes a run's identity and role within an engineering system process.
type RunAutomationDetails struct {
	// A description of the identity and role played within the engineering system by this object's containing run object.
	Description *Message `json:"description,omitempty"`

	// A hierarchical string that uniquely identifies this object's containing run object.
	ID string `json:"id,omitempty"`

	// A stable, unique identifier for this object's containing run object in the form of a GUID.
	GuID string `json:"guid,omitempty"`

	// A stable, unique identifier for the equivalence class of runs to which this object's containing run object belongs in the form of a GUID.
	CorrelationGuID string `json:"correlationGuid,omitempty"`

	// Key/value pairs that provide additional information about the run automation details.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// NewRunAutomationDetails - creates a new
func NewRunAutomationDetails() *RunAutomationDetails {
	return &RunAutomationDetails{}
}

// WithDescription - add a Description to the RunAutomationDetails
func (d *RunAutomationDetails) WithDescription(description *Message) *RunAutomationDetails {
	d.Description = description
	return d
}

// WithID - add a ID to the RunAutomationDetails
func (i *RunAutomationDetails) WithID(id string) *RunAutomationDetails {
	i.ID = id
	return i
}

// WithGuID - add a GuID to the RunAutomationDetails
func (g *RunAutomationDetails) WithGuID(guid string) *RunAutomationDetails {
	g.GuID = guid
	return g
}

// WithCorrelationGuID - add a CorrelationGuID to the RunAutomationDetails
func (c *RunAutomationDetails) WithCorrelationGuID(correlationGuid string) *RunAutomationDetails {
	c.CorrelationGuID = correlationGuid
	return c
}

// WithProperties - add a Properties to the RunAutomationDetails
func (p *RunAutomationDetails) WithProperties(properties *PropertyBag) *RunAutomationDetails {
	p.Properties = properties
	return p
}
