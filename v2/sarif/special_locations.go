package sarif

// SpecialLocations ...
type SpecialLocations struct {
	PropertyBag
	DisplayBase *ArtifactLocation `json:"displayBase,omitempty"`
}

// NewSpecialLocations ...
func NewSpecialLocations() *SpecialLocations {
	return &SpecialLocations{}
}

// WithDisplayBase sets the DisplayBase
func (specialLocations *SpecialLocations) WithDisplayBase(displayBase *ArtifactLocation) *SpecialLocations {
	specialLocations.DisplayBase = displayBase
	return specialLocations
}
