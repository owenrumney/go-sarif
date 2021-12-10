package sarif

// Tool ...
type Tool struct {
	PropertyBag
	Driver *ToolComponent `json:"driver"`
}

// NewTool ...
func NewTool(driver *ToolComponent) *Tool {
	return &Tool{
		Driver: driver,
	}
}

// NewSimpleTool ...
func NewSimpleTool(driverName string) *Tool {
	return &Tool{
		Driver: NewDriver(driverName),
	}
}
