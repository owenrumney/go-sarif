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
	Properties       map[string]string `json:"properties"`
}

func CreateDriver(name, informationUri string) *Driver {
	return &Driver{
		Name:           name,
		InformationUri: informationUri,
	}
}

func (driver *Driver) GetOrCreate(rule *Rule) (int, error) {
	for i, r := range driver.Rules {
		if r.Id == rule.Id {
			return i, nil
		}
	}
	driver.Rules = append(driver.Rules, rule)
	return len(driver.Rules), nil
}

func CreateRule(id, description, helpUri string, properties map[string]string) *Rule {
	return &Rule{
		Id: id,
		ShortDescription: &TextBlock{
			Text: description,
		},
		HelpUri:    helpUri,
		Properties: properties,
	}
}

func CreateTool(driver *Driver) *Tool {
	return &Tool{
		Driver: driver,
	}
}
