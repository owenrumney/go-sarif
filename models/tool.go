package models

type Tool struct {
	Driver *Driver `json:"driver"`
}

type Driver struct {
	Name           string  `json:"name"`
	InformationUri string  `json:"informationUri"`
	Rules          []*Rule `json:"rules,omitempty"`
}

type Rule struct {
	Id               string            `json:"id"`
	ShortDescription *TextBlock        `json:"shortDescription"`
	HelpUri          string            `json:"helpUri"`
	Properties       map[string]string `json:"properties,omitempty"`
}

func (driver *Driver) GetOrCreateRule(rule *Rule) (int, error) {
	for i, r := range driver.Rules {
		if r.Id == rule.Id {
			return i, nil
		}
	}
	driver.Rules = append(driver.Rules, rule)
	return len(driver.Rules) - 1, nil
}

func NewRule(id, description, helpUri string, properties map[string]string) *Rule {
	return &Rule{
		Id: id,
		ShortDescription: &TextBlock{
			Text: description,
		},
		HelpUri:    helpUri,
		Properties: properties,
	}
}

func NewTool(name, informationUri string) *Tool {
	return &Tool{
		Driver: &Driver{
			Name:           name,
			InformationUri: informationUri,
		},
	}
}
