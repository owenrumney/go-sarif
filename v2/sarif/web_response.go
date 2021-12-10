package sarif

type WebResponse struct {
	PropertyBag
	Body               *ArtifactContent  `json:"body,omitempty"`
	Headers            map[string]string `json:"headers,omitempty"`
	Index              *int              `json:"index,omitempty"`
	NoResponseReceived *bool             `json:"noResponseReceived,omitempty"`
	Protocol           *string           `json:"protocol,omitempty"`
	ReasonPhrase       *string           `json:"reasonPhrase,omitempty"`
	StatusCode         *int              `json:"statusCode,omitempty"`
	Version            *string           `json:"version,omitempty"`
}

func NewWebResponse() *WebResponse {
	return &WebResponse{}
}

func (webResponse *WebResponse) WithBody(body *ArtifactContent) *WebResponse {
	webResponse.Body = body
	return webResponse
}

func (webResponse *WebResponse) WithHeaders(headers map[string]string) *WebResponse {
	webResponse.Headers = headers
	return webResponse
}

func (webResponse *WebResponse) SetHeader(name, value string) {
	webResponse.Headers[name] = value
}

func (webResponse *WebResponse) WithIndex(index int) *WebResponse {
	webResponse.Index = &index
	return webResponse
}

func (webResponse *WebResponse) WithNoResponseReceived(noResponse bool) *WebResponse {
	webResponse.NoResponseReceived = &noResponse
	return webResponse
}

func (webResponse *WebResponse) WithProtocol(protocol string) *WebResponse {
	webResponse.Protocol = &protocol
	return webResponse
}

func (webResponse *WebResponse) WithReasonPhrase(reason string) *WebResponse {
	webResponse.ReasonPhrase = &reason
	return webResponse
}
func (webResponse *WebResponse) WithStatusCode(statusCode int) *WebResponse {
	webResponse.StatusCode = &statusCode
	return webResponse
}

func (webResponse *WebResponse) WithVersion(version string) *WebResponse {
	webResponse.Version = &version
	return webResponse
}
