package sarif

import "encoding/json"

// PropertyBag - Key/value pairs that provide additional information about the object.
type PropertyBag struct {
	// AdditionalProperties - additional properties
	Properties Properties `json:"properties,omitempty"`

	// A set of distinct strings that provide additional information.
	Tags []string `json:"tags,omitempty"`
}

// NewPropertyBag - creates a new
func NewPropertyBag() *PropertyBag {
	return &PropertyBag{
		Properties: make(Properties),
		Tags:       make([]string, 0),
	}
}

// AddProperty - add a property to the properties
func (p *PropertyBag) Add(key string, value interface{}) *PropertyBag {
	p.Properties[key] = value
	return p
}

// WithTags - add a Tags to the PropertyBag
func (t *PropertyBag) WithTags(tags []string) *PropertyBag {
	t.Tags = tags
	return t
}

// AddTag - add a single Tag to the PropertyBag
func (t *PropertyBag) AddTag(tag string) *PropertyBag {
	t.Tags = append(t.Tags, tag)
	return t
}

// MarshalJSON - custom JSON marshaller for PropertyBag
func (p PropertyBag) MarshalJSON() ([]byte, error) {
	// type Alias PropertyBag
	aux := make(map[string]interface{})
	for k, v := range p.Properties {
		aux[k] = v
	}
	if len(p.Tags) > 0 {
		aux["tags"] = p.Tags
	}
	return json.Marshal(aux)
}

// UnmarshalJSON - custom JSON unmarshaller for PropertyBag
func (p *PropertyBag) UnmarshalJSON(data []byte) error {
	// type Alias PropertyBag
	aux := struct {
		Tags []string `json:"tags,omitempty"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.Tags = aux.Tags

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	delete(raw, "tags")
	p.Properties = raw
	return nil
}
