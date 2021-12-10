package sarif

type Conversion struct {
	PropertyBag
	AnalysisToolLogFiles []*ArtifactLocation `json:"analysisToolLogFiles,omitempty"`
	Invocation           *Invocation         `json:"invocation,omitempty"`
	Tool                 *Tool               `json:"tool"`
}

func NewConversion() *Conversion {
	return &Conversion{}
}

func (conversion *Conversion) WithInvocation(invocation *Invocation) *Conversion {
	conversion.Invocation = invocation
	return conversion
}

func (conversion *Conversion) WithTool(tool *Tool) *Conversion {
	conversion.Tool = tool
	return conversion
}
