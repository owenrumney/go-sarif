package sarif

type Tool struct {
	PropertyBag
	Driver *ToolComponent `json:"driver"`
}

func NewTool(driver *ToolComponent) *Tool {
	return &Tool{
		Driver: driver,
	}
}

func NewSimpleTool(driverName string) *Tool {
	return &Tool{
		Driver: NewDriver(driverName),
	}
}
