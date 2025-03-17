package sarif

// Artifact - A single artifact. In some cases, this artifact might be nested within another artifact.
type Artifact struct {
	// The role or roles played by the artifact in the analysis.
	Roles []string `json:"roles"`

	// The contents of the artifact.
	Contents *ArtifactContent `json:"contents,omitempty"`

	// Specifies the encoding for an artifact object that refers to a text file.
	Encoding string `json:"encoding,omitempty"`

	// Specifies the source language for any artifact object that refers to a text file that contains source code.
	SourceLanguage string `json:"sourceLanguage,omitempty"`

	// A short description of the artifact.
	Description *Message `json:"description,omitempty"`

	// The location of the artifact.
	Location *ArtifactLocation `json:"location,omitempty"`

	// The MIME type (RFC 2045) of the artifact.
	MimeType string `json:"mimeType,omitempty"`

	// A dictionary, each of whose keys is the name of a hash function and each of whose values is the hashed value of the artifact produced by the specified hash function.
	Hashes map[string]string `json:"hashes,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the artifact was most recently modified. See "Date/time properties" in the SARIF spec for the required format.
	LastModifiedTimeUtc string `json:"lastModifiedTimeUtc,omitempty"`

	// Key/value pairs that provide additional information about the artifact.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Identifies the index of the immediate parent of the artifact, if this artifact is nested.
	ParentIndex int `json:"parentIndex,omitempty"`

	// The offset in bytes of the artifact within its containing artifact.
	Offset int `json:"offset,omitempty"`

	// The length of the artifact in bytes.
	Length int `json:"length,omitempty"`
}

// NewArtifact - creates a new
func NewArtifact() *Artifact {
	return &Artifact{
		Roles: make([]string, 0),
	}
}

// WithRoles - add a Roles to the Artifact
func (r *Artifact) WithRoles(roles []string) *Artifact {
	r.Roles = roles
	return r
}

// AddRole - add a single Role to the Artifact
func (r *Artifact) AddRole(role string) *Artifact {
	r.Roles = append(r.Roles, role)
	return r
}

// WithContents - add a Contents to the Artifact
func (c *Artifact) WithContents(contents *ArtifactContent) *Artifact {
	c.Contents = contents
	return c
}

// WithEncoding - add a Encoding to the Artifact
func (e *Artifact) WithEncoding(encoding string) *Artifact {
	e.Encoding = encoding
	return e
}

// WithSourceLanguage - add a SourceLanguage to the Artifact
func (s *Artifact) WithSourceLanguage(sourceLanguage string) *Artifact {
	s.SourceLanguage = sourceLanguage
	return s
}

// WithDescription - add a Description to the Artifact
func (d *Artifact) WithDescription(description *Message) *Artifact {
	d.Description = description
	return d
}

// WithLocation - add a Location to the Artifact
func (l *Artifact) WithLocation(location *ArtifactLocation) *Artifact {
	l.Location = location
	return l
}

// WithMimeType - add a MimeType to the Artifact
func (m *Artifact) WithMimeType(mimeType string) *Artifact {
	m.MimeType = mimeType
	return m
}

// AddHash - add a single Hash to the Artifact
func (h *Artifact) AddHash(key, hash string) *Artifact {
	h.Hashes[key] = hash
	return h
}

// WithHashes - add a Hashes to the Artifact
func (h *Artifact) WithHashes(hashes map[string]string) *Artifact {
	h.Hashes = hashes
	return h
}

// WithLastModifiedTimeUtc - add a LastModifiedTimeUtc to the Artifact
func (l *Artifact) WithLastModifiedTimeUtc(lastModifiedTimeUtc string) *Artifact {
	l.LastModifiedTimeUtc = lastModifiedTimeUtc
	return l
}

// WithProperties - add a Properties to the Artifact
func (p *Artifact) WithProperties(properties *PropertyBag) *Artifact {
	p.Properties = properties
	return p
}

// WithParentIndex - add a ParentIndex to the Artifact
func (p *Artifact) WithParentIndex(parentIndex int) *Artifact {
	p.ParentIndex = parentIndex
	return p
}

// WithOffset - add a Offset to the Artifact
func (o *Artifact) WithOffset(offset int) *Artifact {
	o.Offset = offset
	return o
}

// WithLength - add a Length to the Artifact
func (l *Artifact) WithLength(length int) *Artifact {
	l.Length = length
	return l
}
