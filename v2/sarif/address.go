package sarif

type Address struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10541049
	PropertyBag
	Index              *uint   `json:"index,omitempty"`
	AbsoluteAddress    *uint   `json:"absoluteAddress,omitempty"`
	RelativeAddress    *int    `json:"relativeAddress,omitempty"`
	OffsetFromParent   *int    `json:"offsetFromParent,omitempty"`
	Length             *int    `json:"length,omitempty"`
	Name               *string `json:"name,omitempty"`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	Kind               *string `json:"kind,omitempty"`
	ParentIndex        *uint   `json:"parentIndex,omitempty"`
}

func NewAddress() *Address {
	return &Address{}
}

func (address *Address) WithIndex(index int) *Address {
	i := uint(index)
	address.Index = &i
	return address
}

func (address *Address) WithAbsoluteAddress(absoluteAddress int) *Address {
	i := uint(absoluteAddress)
	address.AbsoluteAddress = &i
	return address
}

func (address *Address) WithRelativeAddress(relativeAddress int) *Address {
	address.RelativeAddress = &relativeAddress
	return address
}

func (address *Address) WithOffsetFromParent(offsetFromParent int) *Address {
	address.OffsetFromParent = &offsetFromParent
	return address
}

func (address *Address) WithLength(length int) *Address {
	address.Length = &length
	return address
}

func (address *Address) WithName(name string) *Address {
	address.Name = &name
	return address
}

func (address *Address) WithFullyQualifiedName(fullyQualifiedName string) *Address {
	address.FullyQualifiedName = &fullyQualifiedName
	return address
}

func (address *Address) WithKind(kind string) *Address {
	address.Kind = &kind
	return address
}

func (address *Address) WithParentIndex(parentIndex int) *Address {
	i := uint(parentIndex)
	address.ParentIndex = &i
	return address
}
