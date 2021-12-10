package sarif

import "time"

type Notification struct {
	PropertyBag
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`
	Descriptor     *ReportingDescriptorReference `json:"descriptor,omitempty"`
	Exception      *Exception                    `json:"exception,omitempty"`
	Level          string                        `json:"level,omitempty"`
	Locations      []*Location                   `json:"locations,omitempty"`
	Message        *Message                      `json:"message"`
	ThreadID       *int                          `json:"threadId,omitempty"`
	TimeUTC        *time.Time                    `json:"timeUtc,omitempty"`
}

func NewNotification() *Notification {
	return &Notification{}
}

func (notification *Notification) WithAssociatedRule(associatedRule *ReportingDescriptorReference) *Notification {
	notification.AssociatedRule = associatedRule

	return notification
}

func (notification *Notification) WithDescriptor(descriptor *ReportingDescriptorReference) *Notification {
	notification.Descriptor = descriptor

	return notification
}

func (notification *Notification) WithException(exception *Exception) *Notification {
	notification.Exception = exception

	return notification
}

func (notification *Notification) WithLevel(level string) *Notification {
	notification.Level = level

	return notification
}

func (notification *Notification) WithLocations(locations []*Location) *Notification {
	notification.Locations = locations

	return notification
}

func (notification *Notification) AddLocation(location *Location) {
	notification.Locations = append(notification.Locations, location)
}

func (notification *Notification) WithMessage(message *Message) *Notification {
	notification.Message = message
	return notification
}

func (notification *Notification) WithTextMessage(text string) *Notification {
	if notification.Message == nil {
		notification.Message = &Message{}
	}
	notification.Message.Text = &text
	return notification
}

func (notification *Notification) WithMessageMarkdown(markdown string) *Notification {
	if notification.Message == nil {
		notification.Message = &Message{}
	}
	notification.Message.Markdown = &markdown
	return notification
}

func (notification *Notification) WithThreadID(threadID int) *Notification {
	notification.ThreadID = &threadID

	return notification
}

func (notification *Notification) WithTimeUTC(timeUTC *time.Time) *Notification {
	notification.TimeUTC = timeUTC
	return notification
}
