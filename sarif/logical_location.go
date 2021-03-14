package sarif

type LogicalLocation struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Ref493404505
	Index              *uint   `json:"index,omitempty"`
	Name               *string `json:"name,omitempty"`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	DecoratedName      *string `json:"decoratedName,omitempty"`
	Kind               *string `json:"kind,omitempty"`
	ParentIndex        *uint   `json:"parentIndex,omitempty"`
}
