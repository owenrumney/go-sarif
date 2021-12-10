package sarif

type ExternalPropertyFileReferences struct {
	PropertyBag
	Addresses              []*ExternalPropertyFileReference `json:"addresses,omitempty"`
	Artifacts              []*ExternalPropertyFileReference `json:"artifacts,omitempty"`
	Conversion             *ExternalPropertyFileReference   `json:"conversion,omitempty"`
	Driver                 *ExternalPropertyFileReference   `json:"driver,omitempty"`
	Extensions             []*ExternalPropertyFileReference `json:"extensions,omitempty"`
	ExternalizedProperties *ExternalPropertyFileReference   `json:"externalizedProperties,omitempty"`
	Graphs                 []*ExternalPropertyFileReference `json:"graphs,omitempty"`
	Invocations            []*ExternalPropertyFileReference `json:"invocations,omitempty"`
	LogicalLocations       []*ExternalPropertyFileReference `json:"logicalLocations,omitempty"`
	Policies               []*ExternalPropertyFileReference `json:"policies,omitempty"`
	Properties             *PropertyBag                     `json:"properties,omitempty"`
	Results                []*ExternalPropertyFileReference `json:"results,omitempty"`
	Taxonomies             []*ExternalPropertyFileReference `json:"taxonomies,omitempty"`
	ThreadFlowLocations    []*ExternalPropertyFileReference `json:"threadFlowLocations,omitempty"`
	Translations           []*ExternalPropertyFileReference `json:"translations,omitempty"`
	WebRequests            []*ExternalPropertyFileReference `json:"webRequests,omitempty"`
	WebResponses           []*ExternalPropertyFileReference `json:"webResponses,omitempty"`
}

func NewExternalPropertyFileReferences() *ExternalPropertyFileReferences {
	return &ExternalPropertyFileReferences{}
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithAddress(addresses []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Addresses = addresses
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddAddress(address *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Addresses = append(externalPropertyFileReferences.Addresses, address)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithArtifact(artifacts []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Artifacts = artifacts
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddArtifact(artifact *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Artifacts = append(externalPropertyFileReferences.Artifacts, artifact)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithConversion(conversion *ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Conversion = conversion
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithDriver(driver *ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Driver = driver
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithExtensions(extensions []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Extensions = extensions
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddExtension(extension *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Extensions = append(externalPropertyFileReferences.Extensions, extension)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithExternalizedProperties(externalizedProperties *ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.ExternalizedProperties = externalizedProperties
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithGraphs(graphs []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Graphs = graphs
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddGraph(graph *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Graphs = append(externalPropertyFileReferences.Graphs, graph)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithInvocations(invocations []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Invocations = invocations
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddInvocation(invocation *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Invocations = append(externalPropertyFileReferences.Invocations, invocation)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithLogicalLocations(logicalLocations []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.LogicalLocations = logicalLocations
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddLogicalLocation(logicalLocation *ExternalPropertyFileReference) {
	externalPropertyFileReferences.LogicalLocations = append(externalPropertyFileReferences.LogicalLocations, logicalLocation)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithPolicies(policies []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Policies = policies
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddPolicy(policy *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Policies = append(externalPropertyFileReferences.Policies, policy)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithResults(results []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Results = results
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddResult(result *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Results = append(externalPropertyFileReferences.Results, result)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithTaxonomies(taxonomies []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Taxonomies = taxonomies
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddTaxonomie(taxonomy *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Taxonomies = append(externalPropertyFileReferences.Taxonomies, taxonomy)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithThreadFlowLocations(threadFlowLocations []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.ThreadFlowLocations = threadFlowLocations
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddThreadFlowLocations(threadFlowLocation *ExternalPropertyFileReference) {
	externalPropertyFileReferences.ThreadFlowLocations = append(externalPropertyFileReferences.ThreadFlowLocations, threadFlowLocation)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithTranslations(translation []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.Translations = translation
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddTranslation(translation *ExternalPropertyFileReference) {
	externalPropertyFileReferences.Translations = append(externalPropertyFileReferences.Translations, translation)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithWebRequests(webRequests []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.WebRequests = webRequests
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddWebRequest(webRequest *ExternalPropertyFileReference) {
	externalPropertyFileReferences.WebRequests = append(externalPropertyFileReferences.WebRequests, webRequest)
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) WithWebResponses(webResponses []*ExternalPropertyFileReference) *ExternalPropertyFileReferences {
	externalPropertyFileReferences.WebResponses = webResponses
	return externalPropertyFileReferences
}

func (externalPropertyFileReferences *ExternalPropertyFileReferences) AddWebResponse(webResponse *ExternalPropertyFileReference) {
	externalPropertyFileReferences.WebResponses = append(externalPropertyFileReferences.WebResponses, webResponse)
}
