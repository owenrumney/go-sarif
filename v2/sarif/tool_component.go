package sarif

type ToolComponent struct {
	PropertyBag
	Name           string                 `json:"name"`
	Version        *string                `json:"version,omitempty"`
	InformationURI *string                `json:"informationUri"`
	Notifications  []*ReportingDescriptor `json:"notifications,omitempty"`
	Rules          []*ReportingDescriptor `json:"rules,omitempty"`
	Taxa           []*ReportingDescriptor `json:"taxa,omitempty"`
}

func NewDriver(name string) *ToolComponent {
	return &ToolComponent{
		Name: name,
	}
}

func NewVersionedDriver(name, version string) *ToolComponent {
	return &ToolComponent{
		Name:    name,
		Version: &version,
	}
}

// WithVersion specifies tool version, in whatever format it natively provides. Returns updated driver.
func (driver *ToolComponent) WithVersion(version string) *ToolComponent {
	driver.Version = &version
	return driver
}

func (driver *ToolComponent) getOrCreateRule(rule *ReportingDescriptor) uint {
	for i, r := range driver.Rules {
		if r.ID == rule.ID {
			return uint(i)
		}
	}
	driver.Rules = append(driver.Rules, rule)
	return uint(len(driver.Rules) - 1)
}

func (driver *ToolComponent) WithInformationURI(informationURI string) *ToolComponent {
	driver.InformationURI = &informationURI
	return driver
}

func (driver *ToolComponent) WithNotifications(notifications []*ReportingDescriptor) *ToolComponent {
	driver.Notifications = notifications
	return driver
}

func (driver *ToolComponent) AddNotification(notification *ReportingDescriptor) {
	driver.Notifications = append(driver.Notifications, notification)
}

func (driver *ToolComponent) WithNRules(rules []*ReportingDescriptor) *ToolComponent {
	driver.Rules = rules
	return driver
}

func (driver *ToolComponent) AddRule(rule *ReportingDescriptor) {
	driver.Rules = append(driver.Rules, rule)
}

func (driver *ToolComponent) WithTaxa(taxa []*ReportingDescriptor) *ToolComponent {
	driver.Taxa = taxa
	return driver
}

func (driver *ToolComponent) AddTaxa(taxa *ReportingDescriptor) {
	driver.Taxa = append(driver.Taxa, taxa)
}
