package sarif

// Result represents the results block in the sarif report
type Result struct {
	PropertyBag
	Guid                *string                         `json:"guid,omitempty"`
	CorrelationGuid     *string                         `json:"correlationGuid,omitempty"`
	RuleID              *string                         `json:"ruleId,omitempty"`
	RuleIndex           *uint                           `json:"ruleIndex,omitempty"`
	Rule                *ReportingDescriptorReference   `json:"rule,omitempty"`
	Taxa                []*ReportingDescriptorReference `json:"taxa,omitempty"`
	Kind                *string                         `json:"kind,omitempty"`
	Level               *string                         `json:"level,omitempty"`
	Message             Message                         `json:"message"`
	Locations           []*Location                     `json:"locations,omitempty"`
	AnalysisTarget      *ArtifactLocation               `json:"analysisTarget,omitempty"`
	WebRequest          *WebRequest                     `json:"webRequest,omitempty"`
	WebResponse         *WebResponse                    `json:"webResponse,omitempty"`
	Fingerprints        map[string]interface{}          `json:"fingerprints,omitempty"`
	PartialFingerprints map[string]interface{}          `json:"partialFingerprints,omitempty"`
	CodeFlows           []*CodeFlow                     `json:"codeFlows,omitempty"`
	Graphs              []*Graph                        `json:"graphs,omitempty"`
	GraphTraversals     []*GraphTraversal               `json:"graphTraversals,omitempty"`
	Stacks              []*Stack                        `json:"stacks,omitempty"`
	RelatedLocations    []*Location                     `json:"relatedLocations,omitempty"`
	Suppressions        []*Suppression                  `json:"suppressions,omitempty"`
	BaselineState       *string                         `json:"baselineState,omitempty"`
	Rank                *float32                        `json:"rank,omitempty"`
	Attachments         []*Attachment                   `json:"attachments,omitempty"`
	WorkItemUris        []string                        `json:"workItemUris,omitempty"` // can be null
	HostedViewerUri     *string                         `json:"hostedViewerUri,omitempty"`
	Provenance          *ResultProvenance               `json:"provenance,omitempty"`
	Fixes               []*Fix                          `json:"fixes,omitempty"`
	OccurrenceCount     *uint                           `json:"occurrenceCount,omitempty"`
}

func newRuleResult(ruleID string) *Result {
	return &Result{
		RuleID: &ruleID,
	}
}

func (result *Result) WithGuid(guid string) *Result {
	result.Guid = &guid
	return result
}

func (result *Result) WithCorrelationGuid(correlationGuid string) *Result {
	result.CorrelationGuid = &correlationGuid
	return result
}

func (result *Result) WithRuleIndex(ruleIndex int) *Result {
	index := uint(ruleIndex)
	result.RuleIndex = &index
	return result
}

func (result *Result) WithRule(rule *ReportingDescriptorReference) *Result {
	result.Rule = rule
	return result
}

func (result *Result) WithTaxa(taxa []*ReportingDescriptorReference) *Result {
	result.Taxa = taxa
	return result
}

func (result *Result) AddTaxa(taxa *ReportingDescriptorReference) {
	result.Taxa = append(result.Taxa, taxa)
}

func (result *Result) WithKind(kind string) *Result {
	result.Kind = &kind
	return result
}

func (result *Result) WithLevel(level string) *Result {
	result.Level = &level
	return result
}

func (result *Result) WithMessage(message *Message) *Result {
	result.Message = *message
	return result
}

func (result *Result) WithLocations(locations []*Location) *Result {
	result.Locations = locations
	return result
}

func (result *Result) AddLocation(location *Location) {
	result.Locations = append(result.Locations, location)
}

func (result *Result) WithAnalysisTarget(target *ArtifactLocation) *Result {
	result.AnalysisTarget = target
	return result
}

func (result *Result) WithFingerPrints(fingerPrints map[string]interface{}) *Result {
	result.Fingerprints = fingerPrints
	return result
}

func (result *Result) SetFingerPrint(name string, value interface{}) {
	result.Fingerprints[name] = value
}

func (result *Result) WithPartialFingerPrints(fingerPrints map[string]interface{}) *Result {
	result.PartialFingerprints = fingerPrints
	return result
}

func (result *Result) SetPartialFingerPrint(name string, value interface{}) {
	result.PartialFingerprints[name] = value
}

func (result *Result) WithCodeFlows(codeFlows []*CodeFlow) *Result {
	result.CodeFlows = codeFlows
	return result
}

func (result *Result) AddCodeFlow(codeFlow *CodeFlow) {
	result.CodeFlows = append(result.CodeFlows, codeFlow)

}

func (result *Result) WithGraphs(graphs []*Graph) *Result {
	result.Graphs = graphs
	return result
}

func (result *Result) AddGraph(graph *Graph) {
	result.Graphs = append(result.Graphs, graph)

}

func (result *Result) WithGraphTraversal(graphTraversals []*GraphTraversal) *Result {
	result.GraphTraversals = graphTraversals
	return result
}

func (result *Result) AddGraphTraversal(graphTraversal *GraphTraversal) {
	result.GraphTraversals = append(result.GraphTraversals, graphTraversal)

}
func (result *Result) WithStack(stacks []*Stack) *Result {
	result.Stacks = stacks
	return result
}

func (result *Result) AddStack(stack *Stack) {
	result.Stacks = append(result.Stacks, stack)

}

func (result *Result) WithRelatedLocations(locations []*Location) *Result {
	result.RelatedLocations = locations
	return result
}

func (result *Result) AddRelatedLocation(location *Location) *Result {
	result.RelatedLocations = append(result.RelatedLocations, location)
	return result
}

func (result *Result) WithSuppression(suppressions []*Suppression) *Result {
	result.Suppressions = suppressions
	return result
}

func (result *Result) AddSuppression(suppression *Suppression) {
	result.Suppressions = append(result.Suppressions, suppression)

}

func (result *Result) WithBaselineState(state string) *Result {
	result.BaselineState = &state
	return result
}

func (result *Result) WithRank(rank float32) *Result {
	result.Rank = &rank
	return result
}

func (result *Result) WithAttachments(attachments []*Attachment) *Result {
	result.Attachments = attachments
	return result
}

func (result *Result) AddAttachment(attachments *Attachment) {
	result.Attachments = append(result.Attachments, attachments)

}

func (result *Result) WithWorkItemUris(workItemUris []string) *Result {
	result.WorkItemUris = workItemUris
	return result
}

func (result *Result) AddWorkItemUri(workItemUri string) {
	result.WorkItemUris = append(result.WorkItemUris, workItemUri)

}

func (result *Result) WithHostedViewerUri(hostedViewerUri string) *Result {
	result.HostedViewerUri = &hostedViewerUri
	return result
}

func (result *Result) WithFix(fixes []*Fix) *Result {
	result.Fixes = fixes
	return result
}

func (result *Result) AddFix(fix *Fix) {
	result.Fixes = append(result.Fixes, fix)
}

func (result *Result) WithOccurrenceCount(occurrenceCount int) *Result {
	count := uint(occurrenceCount)
	result.OccurrenceCount = &count
	return result
}
