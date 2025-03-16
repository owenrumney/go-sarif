package sarif

// ExternalPropertyFileReference - Contains information that enables a SARIF consumer to locate the external property file that contains the value of an externalized property associated with the run.
type ExternalPropertyFileReference struct {
	// A non-negative integer specifying the number of items contained in the external property file.
	ItemCount int `json:"itemCount,omitempty"`

	// Key/value pairs that provide additional information about the external property file.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The location of the external property file.
	Location *ArtifactLocation `json:"location,omitempty"`

	// A stable, unique identifier for the external property file in the form of a GUID.
	Guid *Guid `json:"guid,omitempty"`
}

// NewExternalPropertyFileReference - creates a new
func NewExternalPropertyFileReference() *ExternalPropertyFileReference {
	return &ExternalPropertyFileReference{}
}

// WithItemCount - add a ItemCount to the ExternalPropertyFileReference
func (i *ExternalPropertyFileReference) WithItemCount(itemCount int) *ExternalPropertyFileReference {
	i.ItemCount = itemCount
	return i
}

// WithProperties - add a Properties to the ExternalPropertyFileReference
func (p *ExternalPropertyFileReference) WithProperties(properties *PropertyBag) *ExternalPropertyFileReference {
	p.Properties = properties
	return p
}

// WithLocation - add a Location to the ExternalPropertyFileReference
func (l *ExternalPropertyFileReference) WithLocation(location *ArtifactLocation) *ExternalPropertyFileReference {
	l.Location = location
	return l
}

// WithGuid - add a Guid to the ExternalPropertyFileReference
func (g *ExternalPropertyFileReference) WithGuid(guid *Guid) *ExternalPropertyFileReference {
	g.Guid = guid
	return g
}
