package sarif

// Result represents the results block in the sarif report
type Result struct {
	Guid            *string                         `json:"guid,omitempty"`
	CorrelationGuid *string                         `json:"correlationGuid,omitempty"`
	RuleID          *string                         `json:"ruleId,omitempty"`
	RuleIndex       *uint                           `json:"ruleIndex,omitempty"`
	Rule            *ReportingDescriptorReference   `json:"rule,omitempty"`
	Taxa            []*ReportingDescriptorReference `json:"taxa,omitempty"`
	Kind            *string                         `json:"kind,omitempty"`
	Level           *string                         `json:"level,omitempty"`
	Message         Message                         `json:"message"`
	Locations       []*Location                     `json:"locations,omitempty"`
	AnalysisTarget  *ArtifactLocation               `json:"analysisTarget,omitempty"`
	// WebRequest			*webRequest						`json:"webRequest,omitempty"`
	// WebResponse			*webResponse					`json:"webResponse,omitempty"`
	Fingerprints        map[string]interface{} `json:"fingerprints,omitempty"`
	PartialFingerprints map[string]interface{} `json:"partialFingerprints,omitempty"`
	// CodeFlows			[]*codeFlows					`json:"codeFlows,omitempty"`
	// Graphs				[]*graphs						`json:"graphs,omitempty"`
	// GraphTraversals		[]*graphTraversals				`json:"graphTraversals,omitempty"`
	// Stacks				[]*stack						`json:"stacks,omitempty"`
	RelatedLocations []*Location    `json:"relatedLocations,omitempty"`
	Suppressions     []*Suppression `json:"suppressions,omitempty"`
	BaselineState    *string        `json:"baselineState,omitempty"`
	Rank             *float32       `json:"rank,omitempty"`
	// Attachments			[]*attachment					`json:"attachments,omitempty"`
	WorkItemUris    []string `json:"workItemUris,omitempty"` // can be null
	HostedViewerUri *string  `json:"hostedViewerUri,omitempty"`
	// Provenance			*resultProvenance				`json:"provenance,omitempty"`
	Fixes           []*Fix `json:"fixes,omitempty"`
	OccurrenceCount *uint  `json:"occurrenceCount,omitempty"`
}

type ReportingDescriptorReference struct {
	Id            *string                 `json:"id,omitempty"`
	Index         *uint                   `json:"index,omitempty"`
	Guid          *string                 `json:"guid,omitempty"`
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`
}

type ToolComponentReference struct {
	Name  *string `json:"name"`
	Index *uint   `json:"index"`
	Guid  *string `json:"guid"`
}

type MultiformatMessageString struct {
	Text     string  `json:"text"`
	Markdown *string `json:"markdown,omitempty"`
}

type Address struct {
	Index              *uint   `json:"index,omitempty"`
	AbsoluteAddress    *uint   `json:"absoluteAddress,omitempty"`
	RelativeAddress    *int    `json:"relativeAddress,omitempty"`
	OffsetFromParent   *int    `json:"offsetFromParent,omitempty"`
	Length             *int    `json:"length,omitempty"`
	Name               *string `json:"name,omitempty"`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	Kind               *string `json:"kind,omitempty"`
	ParentIndex        *uint   `json:"parentIndex,omitempty"`
}

type Suppression struct {
	Kind          string    `json:"kind"`
	Status        *string   `json:"status"`
	Location      *Location `json:"location"`
	Guid          *string   `json:"guid"`
	Justification *string   `json:"justification"`
}

type Fix struct {
	Description     *Message          `json:"description,omitempty"`
	ArtifactChanges []*ArtifactChange `json:"artifactChanges"` //	required
}

type ArtifactChange struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Replacements     []*Replacement   `json:"replacements"` //required
}

type Replacement struct {
	DeletedRegion   Region           `json:"deletedRegion"`
	InsertedContent *ArtifactContent `json:"insertedContent,omitempty"`
}

func newRuleResult(ruleID string) *Result {
	return &Result{
		RuleID: &ruleID,
	}
}
