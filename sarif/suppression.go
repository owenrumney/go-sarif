package sarif

// Suppression ...
type Suppression struct {
	PropertyBag
	Kind          string    `json:"kind"`
	Status        *string   `json:"status"`
	Location      *Location `json:"location"`
	Guid          *string   `json:"guid"`
	Justification *string   `json:"justification"`
}

// NewSuppression ...
func NewSuppression(kind string) *Suppression {
	return &Suppression{
		Kind: kind,
	}
}

// WithStatus ...
func (s *Suppression) WithStatus(status string) *Suppression {
	s.Status = &status
	return s
}

// WithLocation ...
func (s *Suppression) WithLocation(location *Location) *Suppression {
	s.Location = location
	return s
}

// WithGuid ...
func (s *Suppression) WithGuid(guid string) *Suppression {
	s.Guid = &guid
	return s
}

// WithJustifcation ...
func (s *Suppression) WithJustifcation(justification string) *Suppression {
	s.Justification = &justification
	return s
}
