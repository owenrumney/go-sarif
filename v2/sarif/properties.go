package sarif

type Properties map[string]interface{}

type PropertyBag struct {
	Properties Properties `json:"properties,omitempty"`
}

func NewPropertyBag() *PropertyBag {
	return &PropertyBag{
		Properties: Properties{},
	}
}

func (propertyBag *PropertyBag) Add(key string, value interface{}) {
	propertyBag.Properties[key] = value
}

func (propertyBag *PropertyBag) AddString(key, value string) {
	propertyBag.Add(key, value)
}

func (propertyBag *PropertyBag) AddBoolean(key string, value bool) {
	propertyBag.Add(key, value)
}

func (propertyBag *PropertyBag) AddInteger(key string, value int) {
	propertyBag.Add(key, value)
}

// AttachPropertyBag adds a property bag to a rule
func (propertyBag *PropertyBag) AttachPropertyBag(pb *PropertyBag) {
	propertyBag.Properties = pb.Properties
}
