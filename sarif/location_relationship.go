package sarif

type LocationRelationship struct {
	Target      uint     `json:"target"`
	Kinds       []string `json:"kinds,omitempty"`
	Description *Message `json:"description,omitempty"`
}
