package sarif

type Region struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10541123
	StartLine      *int             `json:"startLine,omitempty"`
	StartColumn    *int             `json:"startColumn,omitempty"`
	EndLine        *int             `json:"endLine,omitempty"`
	EndColumn      *int             `json:"endColumn,omitempty"`
	CharOffset     *int             `json:"charOffset,omitempty"`
	CharLength     *int             `json:"charLength,omitempty"`
	ByteOffset     *int             `json:"byteOffset,omitempty"`
	ByteLength     *int             `json:"byteLength,omitempty"`
	Snippet        *ArtifactContent `json:"snippet,omitempty"`
	Message        *Message         `json:"message,omitempty"`
	SourceLanguage *string          `json:"sourceLanguage,omitempty"`
}
