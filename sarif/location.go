package sarif

type Location struct {
	Id               *uint                   `json:"id,omitempty"`
	PhysicalLocation *PhysicalLocation       `json:"physicalLocation,omitempty"`
	LogicalLocations []*LogicalLocation      `json:"logicalLocations,omitempty"`
	Message          *Message                `json:"message,omitempty"`
	Annotations      []*Region               `json:"annotations,omitempty"`
	Relationships    []*LocationRelationship `json:"relationships,omitempty"`
}

type PhysicalLocation struct {
	ArtifactLocation *ArtifactLocation `json:"artifactLocation,omitempty"`
	Region           *Region           `json:"region,omitempty"`
	ContextRegion    *Region           `json:"contextRegion,omitempty"`
	Address          *Address          `json:"address,omitempty"`
}

type LogicalLocation struct {
	Index              *uint   `json:"index,omitempty"`
	Name               *string `json:"name,omitempty"`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	DecoratedName      *string `json:"decoratedName,omitempty"`
	Kind               *string `json:"kind,omitempty"`
	ParentIndex        *uint   `json:"parentIndex,omitempty"`
}

type LocationRelationship struct {
	Target      uint     `json:"target"`
	Kinds       []string `json:"kinds,omitempty"`
	Description *Message `json:"description,omitempty"`
}
