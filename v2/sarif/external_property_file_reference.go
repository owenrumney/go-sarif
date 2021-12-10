package sarif

// ExternalPropertyFileReference ...
type ExternalPropertyFileReference struct {
	PropertyBag
	GUID      *string           `json:"guid,omitempty"`
	ItemCount *int              `json:"itemCount,omitempty"`
	Location  *ArtifactLocation `json:"location,omitempty"`
}

// NewExternalPropertyFileReference ...
func NewExternalPropertyFileReference() *ExternalPropertyFileReference {
	return &ExternalPropertyFileReference{}
}

// WithGUID sets the GUID
func (externalPropertyFileReferences *ExternalPropertyFileReference) WithGUID(guid string) *ExternalPropertyFileReference {
	externalPropertyFileReferences.GUID = &guid
	return externalPropertyFileReferences
}

// WithItemCount sets the ItemCount
func (externalPropertyFileReferences *ExternalPropertyFileReference) WithItemCount(itemCount int) *ExternalPropertyFileReference {
	externalPropertyFileReferences.ItemCount = &itemCount
	return externalPropertyFileReferences
}
