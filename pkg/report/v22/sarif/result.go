package sarif

// Result - A result produced by an analysis tool.
type Result struct {
	// An array of references to taxonomy reporting descriptors that are applicable to the result.
	Taxa []*ReportingDescriptorReference `json:"taxa"`

	// Key/value pairs that provide additional information about the result.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Identifies the artifact that the analysis tool was instructed to scan. This need not be the same as the artifact where the result actually occurred.
	AnalysisTarget *ArtifactLocation `json:"analysisTarget,omitempty"`

	// An array of 'codeFlow' objects relevant to the result.
	CodeFlows []*CodeFlow `json:"codeFlows"`

	// An array of zero or more unique graph objects associated with the result.
	Graphs []*Graph `json:"graphs"`

	// A set of suppressions relevant to this result.
	Suppressions []*Suppression `json:"suppressions"`

	// Information about how and when the result was detected.
	Provenance *ResultProvenance `json:"provenance,omitempty"`

	// The stable, unique identifier of the rule, if any, to which this result is relevant.
	RuleID string `json:"ruleId,omitempty"`

	// A value specifying the severity level of the result.
	Level string `json:"level,omitempty"`

	// An array of one or more unique 'graphTraversal' objects.
	GraphTraversals []*GraphTraversal `json:"graphTraversals"`

	// A set of locations relevant to this result.
	RelatedLocations []*Location `json:"relatedLocations"`

	// An absolute URI at which the result can be viewed.
	HostedViewerURI string `json:"hostedViewerURI,omitempty"`

	// A message that describes the result. The first sentence of the message only will be displayed when visible space is limited.
	Message *Message `json:"message,omitempty"`

	// The index within the tool component rules array of the rule object associated with this result.
	RuleIndex int `json:"ruleIndex,omitempty"`

	// A web request associated with this result.
	WebRequest *WebRequest `json:"webRequest,omitempty"`

	// A reference used to locate the rule descriptor relevant to this result.
	Rule *ReportingDescriptorReference `json:"rule,omitempty"`

	// A value that categorizes results by evaluation state.
	Kind string `json:"kind,omitempty"`

	// The set of locations where the result was detected. Specify only one location unless the problem indicated by the result can only be corrected by making a change at every specified location.
	Locations []*Location `json:"locations"`

	// A positive integer specifying the number of times this logically unique result was observed in this run.
	OccurrenceCount int `json:"occurrenceCount,omitempty"`

	// A set of strings each of which individually defines a stable, unique identity for the result.
	Fingerprints map[string]string `json:"fingerprints,omitempty"`

	// The state of a result relative to a baseline of a previous run.
	BaselineState string `json:"baselineState,omitempty"`

	// The URIs of the work items associated with this result.
	WorkItemURIs []string `json:"workItemURIs"`

	// An array of 'stack' objects relevant to the result.
	Stacks []*Stack `json:"stacks"`

	// An array of 'fix' objects, each of which represents a proposed fix to the problem indicated by the result.
	Fixes []*Fix `json:"fixes"`

	// A stable, unique identifier for the equivalence class of logically identical results to which this result belongs, in the form of a GUID.
	CorrelationGuID string `json:"correlationGuid,omitempty"`

	// A set of artifacts relevant to the result.
	Attachments []*Attachment `json:"attachments"`

	// A web response associated with this result.
	WebResponse *WebResponse `json:"webResponse,omitempty"`

	// A stable, unique identifier for the result in the form of a GUID.
	Guid *Guid `json:"guid,omitempty"`

	// A set of strings that contribute to the stable, unique identity of the result.
	PartialFingerprints map[string]string `json:"partialFingerprints,omitempty"`

	// A number representing the priority or importance of the result.
	Rank float64 `json:"rank,omitempty"`
}

// NewResult - creates a new
func NewResult() *Result {
	return &Result{
		Taxa:             make([]*ReportingDescriptorReference, 0),
		CodeFlows:        make([]*CodeFlow, 0),
		Graphs:           make([]*Graph, 0),
		Suppressions:     make([]*Suppression, 0),
		GraphTraversals:  make([]*GraphTraversal, 0),
		RelatedLocations: make([]*Location, 0),
		Locations:        make([]*Location, 0),
		WorkItemURIs:     make([]string, 0),
		Stacks:           make([]*Stack, 0),
		Fixes:            make([]*Fix, 0),
		Attachments:      make([]*Attachment, 0),
	}
}

// WithTaxa - add a Taxa to the Result
func (t *Result) WithTaxa(taxa []*ReportingDescriptorReference) *Result {
	t.Taxa = taxa
	return t
}

// AddTaxa - add a single Taxa to the Result
func (t *Result) AddTaxa(taxa *ReportingDescriptorReference) *Result {
	t.Taxa = append(t.Taxa, taxa)
	return t
}

// WithProperties - add a Properties to the Result
func (p *Result) WithProperties(properties *PropertyBag) *Result {
	p.Properties = properties
	return p
}

// WithAnalysisTarget - add a AnalysisTarget to the Result
func (a *Result) WithAnalysisTarget(analysisTarget *ArtifactLocation) *Result {
	a.AnalysisTarget = analysisTarget
	return a
}

// WithCodeFlows - add a CodeFlows to the Result
func (c *Result) WithCodeFlows(codeFlows []*CodeFlow) *Result {
	c.CodeFlows = codeFlows
	return c
}

// AddCodeFlow - add a single CodeFlow to the Result
func (c *Result) AddCodeFlow(codeFlow *CodeFlow) *Result {
	c.CodeFlows = append(c.CodeFlows, codeFlow)
	return c
}

// WithGraphs - add a Graphs to the Result
func (g *Result) WithGraphs(graphs []*Graph) *Result {
	g.Graphs = graphs
	return g
}

// AddGraph - add a single Graph to the Result
func (g *Result) AddGraph(graph *Graph) *Result {
	g.Graphs = append(g.Graphs, graph)
	return g
}

// WithSuppressions - add a Suppressions to the Result
func (s *Result) WithSuppressions(suppressions []*Suppression) *Result {
	s.Suppressions = suppressions
	return s
}

// AddSuppression - add a single Suppression to the Result
func (s *Result) AddSuppression(suppression *Suppression) *Result {
	s.Suppressions = append(s.Suppressions, suppression)
	return s
}

// WithProvenance - add a Provenance to the Result
func (p *Result) WithProvenance(provenance *ResultProvenance) *Result {
	p.Provenance = provenance
	return p
}

// WithRuleID - add a RuleID to the Result
func (r *Result) WithRuleID(ruleId string) *Result {
	r.RuleID = ruleId
	return r
}

// WithLevel - add a Level to the Result
func (l *Result) WithLevel(level string) *Result {
	l.Level = level
	return l
}

// WithGraphTraversals - add a GraphTraversals to the Result
func (g *Result) WithGraphTraversals(graphTraversals []*GraphTraversal) *Result {
	g.GraphTraversals = graphTraversals
	return g
}

// AddGraphTraversal - add a single GraphTraversal to the Result
func (g *Result) AddGraphTraversal(graphTraversal *GraphTraversal) *Result {
	g.GraphTraversals = append(g.GraphTraversals, graphTraversal)
	return g
}

// WithRelatedLocations - add a RelatedLocations to the Result
func (r *Result) WithRelatedLocations(relatedLocations []*Location) *Result {
	r.RelatedLocations = relatedLocations
	return r
}

// AddRelatedLocation - add a single RelatedLocation to the Result
func (r *Result) AddRelatedLocation(relatedLocation *Location) *Result {
	r.RelatedLocations = append(r.RelatedLocations, relatedLocation)
	return r
}

// WithHostedViewerURI - add a HostedViewerURI to the Result
func (h *Result) WithHostedViewerURI(hostedViewerURI string) *Result {
	h.HostedViewerURI = hostedViewerURI
	return h
}

// WithMessage - add a Message to the Result
func (m *Result) WithMessage(message *Message) *Result {
	m.Message = message
	return m
}

// WithRuleIndex - add a RuleIndex to the Result
func (r *Result) WithRuleIndex(ruleIndex int) *Result {
	r.RuleIndex = ruleIndex
	return r
}

// WithWebRequest - add a WebRequest to the Result
func (w *Result) WithWebRequest(webRequest *WebRequest) *Result {
	w.WebRequest = webRequest
	return w
}

// WithRule - add a Rule to the Result
func (r *Result) WithRule(rule *ReportingDescriptorReference) *Result {
	r.Rule = rule
	return r
}

// WithKind - add a Kind to the Result
func (k *Result) WithKind(kind string) *Result {
	k.Kind = kind
	return k
}

// WithLocations - add a Locations to the Result
func (l *Result) WithLocations(locations []*Location) *Result {
	l.Locations = locations
	return l
}

// AddLocation - add a single Location to the Result
func (l *Result) AddLocation(location *Location) *Result {
	l.Locations = append(l.Locations, location)
	return l
}

// WithOccurrenceCount - add a OccurrenceCount to the Result
func (o *Result) WithOccurrenceCount(occurrenceCount int) *Result {
	o.OccurrenceCount = occurrenceCount
	return o
}

// AddFingerprint - add a single Fingerprint to the Result
func (f *Result) AddFingerprint(key, fingerprint string) *Result {
	f.Fingerprints[key] = fingerprint
	return f
}

// WithFingerprints - add a Fingerprints to the Result
func (f *Result) WithFingerprints(fingerprints map[string]string) *Result {
	f.Fingerprints = fingerprints
	return f
}

// WithBaselineState - add a BaselineState to the Result
func (b *Result) WithBaselineState(baselineState string) *Result {
	b.BaselineState = baselineState
	return b
}

// WithWorkItemURIs - add a WorkItemURIs to the Result
func (w *Result) WithWorkItemURIs(workItemURIs []string) *Result {
	w.WorkItemURIs = workItemURIs
	return w
}

// AddWorkItemURI - add a single WorkItemURI to the Result
func (w *Result) AddWorkItemURI(workItemURI string) *Result {
	w.WorkItemURIs = append(w.WorkItemURIs, workItemURI)
	return w
}

// WithStacks - add a Stacks to the Result
func (s *Result) WithStacks(stacks []*Stack) *Result {
	s.Stacks = stacks
	return s
}

// AddStack - add a single Stack to the Result
func (s *Result) AddStack(stack *Stack) *Result {
	s.Stacks = append(s.Stacks, stack)
	return s
}

// WithFixes - add a Fixes to the Result
func (f *Result) WithFixes(fixes []*Fix) *Result {
	f.Fixes = fixes
	return f
}

// AddFixe - add a single Fixe to the Result
func (f *Result) AddFixe(fixe *Fix) *Result {
	f.Fixes = append(f.Fixes, fixe)
	return f
}

// WithCorrelationGuID - add a CorrelationGuID to the Result
func (c *Result) WithCorrelationGuID(correlationGuid string) *Result {
	c.CorrelationGuID = correlationGuid
	return c
}

// WithAttachments - add a Attachments to the Result
func (a *Result) WithAttachments(attachments []*Attachment) *Result {
	a.Attachments = attachments
	return a
}

// AddAttachment - add a single Attachment to the Result
func (a *Result) AddAttachment(attachment *Attachment) *Result {
	a.Attachments = append(a.Attachments, attachment)
	return a
}

// WithWebResponse - add a WebResponse to the Result
func (w *Result) WithWebResponse(webResponse *WebResponse) *Result {
	w.WebResponse = webResponse
	return w
}

// WithGuid - add a Guid to the Result
func (g *Result) WithGuid(guid *Guid) *Result {
	g.Guid = guid
	return g
}

// AddPartialFingerprint - add a single PartialFingerprint to the Result
func (p *Result) AddPartialFingerprint(key, partialFingerprint string) *Result {
	p.PartialFingerprints[key] = partialFingerprint
	return p
}

// WithPartialFingerprints - add a PartialFingerprints to the Result
func (p *Result) WithPartialFingerprints(partialFingerprints map[string]string) *Result {
	p.PartialFingerprints = partialFingerprints
	return p
}

// WithRank - add a Rank to the Result
func (r *Result) WithRank(rank float64) *Result {
	r.Rank = rank
	return r
}
