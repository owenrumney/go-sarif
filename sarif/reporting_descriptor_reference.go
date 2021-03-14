package sarif


type ReportingDescriptorReference struct {
	Id            *string                 `json:"id,omitempty"`
	Index         *uint                   `json:"index,omitempty"`
	Guid          *string                 `json:"guid,omitempty"`
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`
}
