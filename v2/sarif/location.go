package sarif

type Location struct {
	PropertyBag
	Id               *uint                   `json:"id,omitempty"`
	PhysicalLocation *PhysicalLocation       `json:"physicalLocation,omitempty"`
	LogicalLocations []*LogicalLocation      `json:"logicalLocations,omitempty"`
	Message          *Message                `json:"message,omitempty"`
	Annotations      []*Region               `json:"annotations,omitempty"`
	Relationships    []*LocationRelationship `json:"relationships,omitempty"`
}

func NewLocation() *Location {
	return &Location{}
}

func NewLocationWithPhysicalLocation(physicalLocation *PhysicalLocation) *Location {
	return NewLocation().WithPhysicalLocation(physicalLocation)
}

func (location *Location) WithId(id int) *Location {
	i := uint(id)
	location.Id = &i
	return location
}

func (location *Location) WithPhysicalLocation(physicalLocation *PhysicalLocation) *Location {
	location.PhysicalLocation = physicalLocation
	return location
}

func (location *Location) WithLogicalLocations(logicalLocations []*LogicalLocation) *Location {
	location.LogicalLocations = logicalLocations
	return location
}

func (location *Location) AddLogicalLocations(logicalLocation *LogicalLocation) {
	location.LogicalLocations = append(location.LogicalLocations, logicalLocation)
}

func (location *Location) WithMessage(message *Message) *Location {
	location.Message = message
	return location
}

func (location *Location) WithDescriptionText(text string) *Location {
	if location.Message == nil {
		location.Message = &Message{}
	}
	location.Message.Text = &text
	return location
}

func (location *Location) WithDescriptionMarkdown(markdown string) *Location {
	if location.Message == nil {
		location.Message = &Message{}
	}
	location.Message.Markdown = &markdown
	return location
}

func (location *Location) WithAnnotations(annotations []*Region) *Location {
	location.Annotations = annotations
	return location
}

func (location *Location) AddAnnotation(annotation *Region) {
	location.Annotations = append(location.Annotations, annotation)
}

func (location *Location) WithRelationships(locationRelationships []*LocationRelationship) *Location {
	location.Relationships = locationRelationships
	return location
}

func (location *Location) AddRelationship(locationRelationship *LocationRelationship) {
	location.Relationships = append(location.Relationships, locationRelationship)
}
