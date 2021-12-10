package sarif

type LocationRelationship struct {
	PropertyBag
	Target      uint     `json:"target"`
	Kinds       []string `json:"kinds,omitempty"`
	Description *Message `json:"description,omitempty"`
}

func NewLocationRelationship(target int) *LocationRelationship {
	t := uint(target)
	return &LocationRelationship{
		Target: t,
	}
}

func (locationRelationship *LocationRelationship) WithKinds(kinds []string) *LocationRelationship {
	locationRelationship.Kinds = kinds
	return locationRelationship
}

func (locationRelationship *LocationRelationship) AddKind(kind string) {
	locationRelationship.Kinds = append(locationRelationship.Kinds, kind)
}

func (locationRelationship *LocationRelationship) WithDescription(message *Message) *LocationRelationship {
	locationRelationship.Description = message
	return locationRelationship
}

func (locationRelationship *LocationRelationship) WithDescriptionText(text string) *LocationRelationship {
	if locationRelationship.Description == nil {
		locationRelationship.Description = &Message{}
	}
	locationRelationship.Description.Text = &text
	return locationRelationship
}

func (locationRelationship *LocationRelationship) WithDescriptionMarkdown(markdown string) *LocationRelationship {
	if locationRelationship.Description == nil {
		locationRelationship.Description = &Message{}
	}
	locationRelationship.Description.Markdown = &markdown
	return locationRelationship
}
