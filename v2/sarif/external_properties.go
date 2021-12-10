package sarif

type ExternalProperties struct {
	PropertyBag
	Addresses              []*Address            `json:"addresses,omitempty"`
	Artifacts              []*Artifact           `json:"artifacts,omitempty"`
	Conversion             *Conversion           `json:"conversion,omitempty"`
	Driver                 *ToolComponent        `json:"driver,omitempty"`
	Extensions             []*ToolComponent      `json:"extensions,omitempty"`
	ExternalizedProperties *PropertyBag          `json:"externalizedProperties,omitempty"`
	Graphs                 []*Graph              `json:"graphs,omitempty"`
	GUID                   *string               `json:"guid,omitempty"`
	Invocations            []*Invocation         `json:"invocations,omitempty"`
	LogicalLocations       []*LogicalLocation    `json:"logicalLocations,omitempty"`
	Policies               []*ToolComponent      `json:"policies,omitempty"`
	Results                []*Result             `json:"results,omitempty"`
	RunGUID                *string               `json:"runGuid,omitempty"`
	Schema                 *string               `json:"schema,omitempty"`
	Taxonomies             []*ToolComponent      `json:"taxonomies,omitempty"`
	ThreadFlowLocations    []*ThreadFlowLocation `json:"threadFlowLocations,omitempty"`
	Translations           []*ToolComponent      `json:"translations,omitempty"`
	Version                string                `json:"version,omitempty"`
	WebRequests            []*WebRequest         `json:"webRequests,omitempty"`
	WebResponses           []*WebResponse        `json:"webResponses,omitempty"`
}

func NewExternalProperties() *ExternalProperties {
	return &ExternalProperties{}
}

func (externalProperties *ExternalProperties) WithAddress(addresses []*Address) *ExternalProperties {
	externalProperties.Addresses = addresses
	return externalProperties
}

func (externalProperties *ExternalProperties) AddAddress(address *Address) {
	externalProperties.Addresses = append(externalProperties.Addresses, address)
}

func (externalProperties *ExternalProperties) WithArtifact(artifacts []*Artifact) *ExternalProperties {
	externalProperties.Artifacts = artifacts
	return externalProperties
}

func (externalProperties *ExternalProperties) AddArtifact(artifact *Artifact) {
	externalProperties.Artifacts = append(externalProperties.Artifacts, artifact)
}

func (externalProperties *ExternalProperties) WithConversion(conversion *Conversion) *ExternalProperties {
	externalProperties.Conversion = conversion
	return externalProperties
}

func (externalProperties *ExternalProperties) WithDriver(driver *ToolComponent) *ExternalProperties {
	externalProperties.Driver = driver
	return externalProperties
}

func (externalProperties *ExternalProperties) WithExtensions(extensions []*ToolComponent) *ExternalProperties {
	externalProperties.Extensions = extensions
	return externalProperties
}

func (externalProperties *ExternalProperties) AddExtension(extension *ToolComponent) {
	externalProperties.Extensions = append(externalProperties.Extensions, extension)
}

func (externalProperties *ExternalProperties) WithExternalizedProperties(externalizedProperties *PropertyBag) *ExternalProperties {
	externalProperties.ExternalizedProperties = externalizedProperties
	return externalProperties
}

func (externalProperties *ExternalProperties) WithGraphs(graphs []*Graph) *ExternalProperties {
	externalProperties.Graphs = graphs
	return externalProperties
}

func (externalProperties *ExternalProperties) AddGraph(graph *Graph) {
	externalProperties.Graphs = append(externalProperties.Graphs, graph)
}

func (externalProperties *ExternalProperties) WithGUID(guid string) *ExternalProperties {
	externalProperties.GUID = &guid
	return externalProperties
}

func (externalProperties *ExternalProperties) WithInvocations(invocations []*Invocation) *ExternalProperties {
	externalProperties.Invocations = invocations
	return externalProperties
}

func (externalProperties *ExternalProperties) AddInvocation(invocation *Invocation) {
	externalProperties.Invocations = append(externalProperties.Invocations, invocation)
}

func (externalProperties *ExternalProperties) WithLogicalLocations(logicalLocations []*LogicalLocation) *ExternalProperties {
	externalProperties.LogicalLocations = logicalLocations
	return externalProperties
}

func (externalProperties *ExternalProperties) AddLogicalLocation(logicalLocation *LogicalLocation) {
	externalProperties.LogicalLocations = append(externalProperties.LogicalLocations, logicalLocation)
}

func (externalProperties *ExternalProperties) WithPolicies(policies []*ToolComponent) *ExternalProperties {
	externalProperties.Policies = policies
	return externalProperties
}

func (externalProperties *ExternalProperties) AddPolicy(policy *ToolComponent) {
	externalProperties.Policies = append(externalProperties.Policies, policy)
}

func (externalProperties *ExternalProperties) WithResults(results []*Result) *ExternalProperties {
	externalProperties.Results = results
	return externalProperties
}

func (externalProperties *ExternalProperties) AddResult(result *Result) {
	externalProperties.Results = append(externalProperties.Results, result)
}

func (externalProperties *ExternalProperties) WithRunGUID(runGUID string) *ExternalProperties {
	externalProperties.RunGUID = &runGUID
	return externalProperties
}

func (externalProperties *ExternalProperties) WithSchema(schema string) *ExternalProperties {
	externalProperties.Schema = &schema
	return externalProperties
}

func (externalProperties *ExternalProperties) WithVersion(version string) *ExternalProperties {
	externalProperties.Version = version
	return externalProperties
}

func (externalProperties *ExternalProperties) WithTaxonomies(taxonomies []*ToolComponent) *ExternalProperties {
	externalProperties.Taxonomies = taxonomies
	return externalProperties
}

func (externalProperties *ExternalProperties) AddTaxonomie(taxonomy *ToolComponent) {
	externalProperties.Taxonomies = append(externalProperties.Taxonomies, taxonomy)
}

func (externalProperties *ExternalProperties) WithThreadFlowLocations(threadFlowLocations []*ThreadFlowLocation) *ExternalProperties {
	externalProperties.ThreadFlowLocations = threadFlowLocations
	return externalProperties
}

func (externalProperties *ExternalProperties) AddThreadFlowLocations(threadFlowLocation *ThreadFlowLocation) {
	externalProperties.ThreadFlowLocations = append(externalProperties.ThreadFlowLocations, threadFlowLocation)
}

func (externalProperties *ExternalProperties) WithTranslations(translation []*ToolComponent) *ExternalProperties {
	externalProperties.Translations = translation
	return externalProperties
}

func (externalProperties *ExternalProperties) AddTranslation(translation *ToolComponent) {
	externalProperties.Translations = append(externalProperties.Translations, translation)
}

func (externalProperties *ExternalProperties) WithWebRequests(webRequests []*WebRequest) *ExternalProperties {
	externalProperties.WebRequests = webRequests
	return externalProperties
}

func (externalProperties *ExternalProperties) AddWebRequest(webRequest *WebRequest) {
	externalProperties.WebRequests = append(externalProperties.WebRequests, webRequest)
}

func (externalProperties *ExternalProperties) WithWebResponses(webResponses []*WebResponse) *ExternalProperties {
	externalProperties.WebResponses = webResponses
	return externalProperties
}

func (externalProperties *ExternalProperties) AddWebResponse(webResponse *WebResponse) {
	externalProperties.WebResponses = append(externalProperties.WebResponses, webResponse)
}
