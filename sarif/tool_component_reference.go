package sarif

type ToolComponentReference struct {
	Name  *string `json:"name"`
	Index *uint   `json:"index"`
	Guid  *string `json:"guid"`
}
