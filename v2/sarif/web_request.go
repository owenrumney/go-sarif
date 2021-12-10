package sarif

type WebRequest struct {
	PropertyBag
	Body       *ArtifactContent  `json:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Index      *int              `json:"index,omitempty"`
	Method     *string           `json:"method,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	Protocol   *string           `json:"protocol,omitempty"`
	Target     *string           `json:"target,omitempty"`
	Version    *string           `json:"version,omitempty"`
}

func NewWebRequest() *WebRequest {
	return &WebRequest{}
}

func (webRequest *WebRequest) WithBody(body *ArtifactContent) *WebRequest {
	webRequest.Body = body
	return webRequest
}

func (webRequest *WebRequest) WithHeaders(headers map[string]string) *WebRequest {
	webRequest.Headers = headers
	return webRequest
}

func (webRequest *WebRequest) SetHeader(name, value string) {
	webRequest.Headers[name] = value
}

func (webRequest *WebRequest) WithIndex(index int) *WebRequest {
	webRequest.Index = &index
	return webRequest
}

func (webRequest *WebRequest) WithMethod(method string) *WebRequest {
	webRequest.Method = &method
	return webRequest
}

func (webRequest *WebRequest) WithParameters(parameters map[string]string) *WebRequest {
	webRequest.Parameters = parameters
	return webRequest
}

func (webRequest *WebRequest) SetParameter(name, value string) {
	webRequest.Parameters[name] = value
}

func (webRequest *WebRequest) WithProtocol(protocol string) *WebRequest {
	webRequest.Protocol = &protocol
	return webRequest
}

func (webRequest *WebRequest) WithTarget(target string) *WebRequest {
	webRequest.Target = &target
	return webRequest
}

func (webRequest *WebRequest) WithVersion(version string) *WebRequest {
	webRequest.Version = &version
	return webRequest
}
