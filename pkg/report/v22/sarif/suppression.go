package sarif

// Suppression - A suppression that is relevant to a result.
type Suppression struct {
	// Key/value pairs that provide additional information about the suppression.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A stable, unique identifier for the suprression in the form of a GUID.
	Guid *Guid `json:"guid,omitempty"`

	// A string that indicates where the suppression is persisted.
	Kind string `json:"kind,omitempty"`

	// A string that indicates the review status of the suppression.
	Status string `json:"status,omitempty"`

	// A string representing the justification for the suppression.
	Justification string `json:"justification,omitempty"`

	// Identifies the location associated with the suppression.
	Location *Location `json:"location,omitempty"`
}

// NewSuppression - creates a new
func NewSuppression() *Suppression {
	return &Suppression{}
}

// WithProperties - add a Properties to the Suppression
func (p *Suppression) WithProperties(properties *PropertyBag) *Suppression {
	p.Properties = properties
	return p
}

// WithGuid - add a Guid to the Suppression
func (g *Suppression) WithGuid(guid *Guid) *Suppression {
	g.Guid = guid
	return g
}

// WithKind - add a Kind to the Suppression
func (k *Suppression) WithKind(kind string) *Suppression {
	k.Kind = kind
	return k
}

// WithStatus - add a Status to the Suppression
func (s *Suppression) WithStatus(status string) *Suppression {
	s.Status = status
	return s
}

// WithJustification - add a Justification to the Suppression
func (j *Suppression) WithJustification(justification string) *Suppression {
	j.Justification = justification
	return j
}

// WithLocation - add a Location to the Suppression
func (l *Suppression) WithLocation(location *Location) *Suppression {
	l.Location = location
	return l
}
