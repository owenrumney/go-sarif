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

type ArtifactLocation struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540865
	URI         *string  `json:"uri,omitempty"`
	URIBaseId   *string  `json:"uriBaseId,omitempty"`
	Index       *uint    `json:"index,omitempty"`
	Description *Message `json:"description,omitempty"`
}

type ArtifactContent struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540860
	Text     *string                   `json:"text,omitempty"`
	Binary   *string                   `json:"binary,omitempty"`
	Rendered *MultiformatMessageString `json:"rendered,omitempty"`
}
