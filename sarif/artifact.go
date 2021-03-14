package sarif

type Artifact struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10541049
	Location            *ArtifactLocation `json:"location,omitempty"`
	ParentIndex         *uint             `json:"parentIndex,omitempty"`
	Offset              *uint             `json:"offset,omitempty"`
	Length              int               `json:"length"`
	Roles               []string          `json:"roles,omitempty"`
	MimeType            *string           `json:"mimeType,omitempty"`
	Contents            *ArtifactContent  `json:"contents,omitempty"`
	Encoding            *string           `json:"encoding,omitempty"`
	SourceLanguage      *string           `json:"sourceLanguage,omitempty"`
	Hashes              map[string]string `json:"hashes,omitempty"`
	LastModifiedTimeUtc *string           `json:"lastModifiedTimeUtc,omitempty"`
	Description         *Message          `json:"description,omitempty"`
}
