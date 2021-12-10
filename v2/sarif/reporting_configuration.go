package sarif

type ReportingConfiguration struct {
	PropertyBag
	Enabled    *bool        `json:"enabled,omitempty"`
	Level      string       `json:"level,omitempty"`
	Parameters *PropertyBag `json:"parameters,omitempty"`
	Rank       *float64     `json:"rank,omitempty"`
}

func NewReportingConfiguration() *ReportingConfiguration {
	return &ReportingConfiguration{}
}

func (reportingConfiguration *ReportingConfiguration) WithEnabled(enabled bool) *ReportingConfiguration {
	reportingConfiguration.Enabled = &enabled
	return reportingConfiguration
}

func (reportingConfiguration *ReportingConfiguration) WithLevel(level string) *ReportingConfiguration {
	reportingConfiguration.Level = level
	return reportingConfiguration
}

func (reportingConfiguration *ReportingConfiguration) WithParameters(parameters *PropertyBag) *ReportingConfiguration {
	reportingConfiguration.Parameters = parameters
	return reportingConfiguration
}

func (reportingConfiguration *ReportingConfiguration) WithRank(rank float64) *ReportingConfiguration {
	reportingConfiguration.Rank = &rank
	return reportingConfiguration
}
