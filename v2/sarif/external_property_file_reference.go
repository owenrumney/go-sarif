package sarif

type ExternalPropertyFileReference struct {
	PropertyBag
	GUID      *string           `json:"guid,omitempty"`
	ItemCount *int              `json:"itemCount,omitempty"`
	Location  *ArtifactLocation `json:"location,omitempty"`
}

func NewExternalPropertyFileReference() *ExternalPropertyFileReference {
	return &ExternalPropertyFileReference{}
}

func (externalPropertyFileReferences *ExternalPropertyFileReference) WithGUID(guid string) *ExternalPropertyFileReference {
	externalPropertyFileReferences.GUID = &guid
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReference) WithItemCount(itemCount int) *ExternalPropertyFileReference {
	externalPropertyFileReferences.ItemCount = &itemCount
	return externalPropertyFileReferences
}
