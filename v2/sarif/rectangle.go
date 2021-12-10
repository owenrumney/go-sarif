package sarif

type Rectangle struct {
	PropertyBag
	Bottom  *float64 `json:"bottom,omitempty"`
	Left    *float64 `json:"left,omitempty"`
	Message *Message `json:"message,omitempty"`
	Right   *float64 `json:"right,omitempty"`
	Top     *float64 `json:"top,omitempty"`
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (rectangle *Rectangle) WithBottom(bottom float64) *Rectangle {
	rectangle.Bottom = &bottom
	return rectangle
}

func (rectangle *Rectangle) WithTop(top float64) *Rectangle {
	rectangle.Top = &top
	return rectangle
}

func (rectangle *Rectangle) WithLeft(withLeft float64) *Rectangle {
	rectangle.Left = &withLeft
	return rectangle
}

func (rectangle *Rectangle) WithRight(right float64) *Rectangle {
	rectangle.Right = &right
	return rectangle
}

func (rectangle *Rectangle) WithMessage(message *Message) *Rectangle {
	rectangle.Message = message
	return rectangle
}

func (rectangle *Rectangle) WithTextMessage(text string) *Rectangle {
	if rectangle.Message == nil {
		rectangle.Message = &Message{}
	}
	rectangle.Message.Text = &text
	return rectangle
}

func (rectangle *Rectangle) WithMessageMarkdown(markdown string) *Rectangle {
	if rectangle.Message == nil {
		rectangle.Message = &Message{}
	}
	rectangle.Message.Markdown = &markdown
	return rectangle
}