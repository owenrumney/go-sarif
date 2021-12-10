package sarif

type TranslationMetadata struct {
	PropertyBag
	DownloadURI      *string                   `json:"downloadUri,omitempty"`
	FullDescription  *MultiformatMessageString `json:"fullDescription,omitempty"`
	FullName         *string                   `json:"fullName,omitempty"`
	InformationURI   *string                   `json:"informationUri,omitempty"`
	Name             *string                   `json:"name"`
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
}

func NewTranslationMetadata() *TranslationMetadata {
	return &TranslationMetadata{}
}

func (translationMetadata *TranslationMetadata) WithDownloadURI(downloadURI string) *TranslationMetadata {
	translationMetadata.DownloadURI = &downloadURI
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithFullDescription(message *MultiformatMessageString) *TranslationMetadata {
	translationMetadata.FullDescription = message
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithFullDescriptionText(text string) *TranslationMetadata {
	if translationMetadata.FullDescription == nil {
		translationMetadata.FullDescription = &MultiformatMessageString{}
	}
	translationMetadata.FullDescription.Text = &text
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithFullDescriptionMarkdown(markdown string) *TranslationMetadata {
	if translationMetadata.FullDescription == nil {
		translationMetadata.FullDescription = &MultiformatMessageString{}
	}
	translationMetadata.FullDescription.Markdown = &markdown
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithFullName(fullname string) *TranslationMetadata {
	translationMetadata.FullName = &fullname
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithInformationURI(informationURI string) *TranslationMetadata {
	translationMetadata.InformationURI = &informationURI
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithName(name string) *TranslationMetadata {
	translationMetadata.Name = &name

	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithShortDescription(message *MultiformatMessageString) *TranslationMetadata {
	translationMetadata.ShortDescription = message
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithShortShortDescriptionText(text string) *TranslationMetadata {
	if translationMetadata.ShortDescription == nil {
		translationMetadata.ShortDescription = &MultiformatMessageString{}
	}
	translationMetadata.ShortDescription.Text = &text
	return translationMetadata
}

func (translationMetadata *TranslationMetadata) WithShortDescriptionMarkdown(markdown string) *TranslationMetadata {
	if translationMetadata.ShortDescription == nil {
		translationMetadata.ShortDescription = &MultiformatMessageString{}
	}
	translationMetadata.ShortDescription.Markdown = &markdown
	return translationMetadata
}
