package sarif

type Location struct {
	Id               *uint                   `json:"id,omitempty"`
	PhysicalLocation *PhysicalLocation       `json:"physicalLocation,omitempty"`
	LogicalLocations []*LogicalLocation      `json:"logicalLocations,omitempty"`
	Message          *Message                `json:"message,omitempty"`
	Annotations      []*Region               `json:"annotations,omitempty"`
	Relationships    []*LocationRelationship `json:"relationships,omitempty"`
}
