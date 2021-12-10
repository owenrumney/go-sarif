package sarif

type ConfigurationOverride struct {
	PropertyBag
	Configuration *ReportingConfiguration       `json:"configuration,omitempty"`
	Descriptor    *ReportingDescriptorReference `json:"descriptor,omitempty"`
}

// NewConfigurationOverride creates a new ConfigurationOverride and returns a pointer to it
func NewConfigurationOverride() *ConfigurationOverride {
	return &ConfigurationOverride{}
}

func (configurationOverride *ConfigurationOverride) WithDescriptor(descriptor *ReportingDescriptorReference) *ConfigurationOverride {
	configurationOverride.Descriptor = descriptor
	return configurationOverride
}

func (configurationOverride *ConfigurationOverride) WithConfiguration(configuration *ReportingConfiguration) *ConfigurationOverride {
	configurationOverride.Configuration = configuration
	return configurationOverride
}
