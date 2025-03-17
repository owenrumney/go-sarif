package sarif

// WebResponse - Describes the response to an HTTP request.
type WebResponse struct {
	// The response protocol. Example: 'http'.
	Protocol string `json:"protocol,omitempty"`

	// The response version. Example: '1.1'.
	Version string `json:"version,omitempty"`

	// The response reason. Example: 'Not found'.
	ReasonPhrase string `json:"reasonPhrase,omitempty"`

	// The body of the response.
	Body *ArtifactContent `json:"body,omitempty"`

	// Specifies whether a response was received from the server.
	NoResponseReceived int `json:"noResponseReceived,omitempty"`

	// Key/value pairs that provide additional information about the response.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The index within the run.webResponses array of the response object associated with this result.
	Index int `json:"index,omitempty"`

	// The response status code. Example: 451.
	StatusCode int `json:"statusCode,omitempty"`

	// The response headers.
	Headers map[string]string `json:"headers,omitempty"`
}

// NewWebResponse - creates a new
func NewWebResponse() *WebResponse {
	return &WebResponse{}
}

// WithProtocol - add a Protocol to the WebResponse
func (p *WebResponse) WithProtocol(protocol string) *WebResponse {
	p.Protocol = protocol
	return p
}

// WithVersion - add a Version to the WebResponse
func (v *WebResponse) WithVersion(version string) *WebResponse {
	v.Version = version
	return v
}

// WithReasonPhrase - add a ReasonPhrase to the WebResponse
func (r *WebResponse) WithReasonPhrase(reasonPhrase string) *WebResponse {
	r.ReasonPhrase = reasonPhrase
	return r
}

// WithBody - add a Body to the WebResponse
func (b *WebResponse) WithBody(body *ArtifactContent) *WebResponse {
	b.Body = body
	return b
}

// WithNoResponseReceived - add a NoResponseReceived to the WebResponse
func (n *WebResponse) WithNoResponseReceived(noResponseReceived int) *WebResponse {
	n.NoResponseReceived = noResponseReceived
	return n
}

// WithProperties - add a Properties to the WebResponse
func (p *WebResponse) WithProperties(properties *PropertyBag) *WebResponse {
	p.Properties = properties
	return p
}

// WithIndex - add a Index to the WebResponse
func (i *WebResponse) WithIndex(index int) *WebResponse {
	i.Index = index
	return i
}

// WithStatusCode - add a StatusCode to the WebResponse
func (s *WebResponse) WithStatusCode(statusCode int) *WebResponse {
	s.StatusCode = statusCode
	return s
}

// AddHeader - add a single Header to the WebResponse
func (h *WebResponse) AddHeader(key, header string) *WebResponse {
	h.Headers[key] = header
	return h
}

// WithHeaders - add a Headers to the WebResponse
func (h *WebResponse) WithHeaders(headers map[string]string) *WebResponse {
	h.Headers = headers
	return h
}
