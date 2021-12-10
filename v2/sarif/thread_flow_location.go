package sarif

import "time"

type ThreadFlowLocation struct {
	PropertyBag
	ExecutionOrder   *int                                 `json:"executionOrder,omitempty"`
	ExecutionTimeUTC *time.Time                           `json:"executionTimeUtc,omitempty"`
	Importance       interface{}                          `json:"importance,omitempty"`
	Index            *int                                 `json:"index,omitempty"`
	Kinds            []string                             `json:"kinds,omitempty"`
	Location         *Location                            `json:"location,omitempty"`
	Module           *string                              `json:"module,omitempty"`
	NestingLevel     *int                                 `json:"nestingLevel,omitempty"`
	Stack            *Stack                               `json:"stack,omitempty"`
	State            map[string]*MultiformatMessageString `json:"state,omitempty"`
	Taxa             []*ReportingDescriptorReference      `json:"taxa,omitempty"`
	WebRequest       *WebRequest                          `json:"webRequest,omitempty"`
	WebResponse      *WebResponse                         `json:"webResponse,omitempty"`
}

func NewThreadFlowLocation() *ThreadFlowLocation {
	return &ThreadFlowLocation{}
}

func (threadFlowLocation *ThreadFlowLocation) WithExecutionOrder(order int) *ThreadFlowLocation {
	threadFlowLocation.ExecutionOrder = &order
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithExecutionTimeUTC(executionTimeUTC *time.Time) *ThreadFlowLocation {
	threadFlowLocation.ExecutionTimeUTC = executionTimeUTC
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithImportance(importance interface{}) *ThreadFlowLocation {
	threadFlowLocation.Importance = importance
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithIndex(index int) *ThreadFlowLocation {
	threadFlowLocation.Index = &index
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithKinds(kinds []string) *ThreadFlowLocation {
	threadFlowLocation.Kinds = kinds
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) AddKind(kind string) {
	threadFlowLocation.Kinds = append(threadFlowLocation.Kinds, kind)
}

func (threadFlowLocation *ThreadFlowLocation) WithLocation(location *Location) *ThreadFlowLocation {
	threadFlowLocation.Location = location
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithModule(module string) *ThreadFlowLocation {
	threadFlowLocation.Module = &module
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithNestingLevel(nestingLevel int) *ThreadFlowLocation {
	threadFlowLocation.NestingLevel = &nestingLevel
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithStack(stack *Stack) *ThreadFlowLocation {
	threadFlowLocation.Stack = stack
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithState(state map[string]*MultiformatMessageString) *ThreadFlowLocation {
	threadFlowLocation.State = state
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithTaxa(taxa []*ReportingDescriptorReference) *ThreadFlowLocation {
	threadFlowLocation.Taxa = taxa
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) AddTaxa(taxa *ReportingDescriptorReference) {
	threadFlowLocation.Taxa = append(threadFlowLocation.Taxa, taxa)
}

func (threadFlowLocation *ThreadFlowLocation) WithWebRequest(webRequest *WebRequest) *ThreadFlowLocation {
	threadFlowLocation.WebRequest = webRequest
	return threadFlowLocation
}

func (threadFlowLocation *ThreadFlowLocation) WithWebResponse(webResponse *WebResponse) *ThreadFlowLocation {
	threadFlowLocation.WebResponse = webResponse
	return threadFlowLocation
}
