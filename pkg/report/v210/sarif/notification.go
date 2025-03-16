package sarif

// Notification - Describes a condition relevant to the tool itself, as opposed to being relevant to a target being analyzed by the tool.
type Notification struct {
	// The thread identifier of the code that generated the notification.
	ThreadID int `json:"threadId,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the analysis tool generated the notification.
	TimeUtc string `json:"timeUtc,omitempty"`

	// The runtime exception, if any, relevant to this notification.
	Exception *Exception `json:"exception,omitempty"`

	// Key/value pairs that provide additional information about the notification.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A message that describes the condition that was encountered.
	Message *Message `json:"message,omitempty"`

	// A reference used to locate the descriptor relevant to this notification.
	Descriptor *ReportingDescriptorReference `json:"descriptor,omitempty"`

	// A reference used to locate the rule descriptor associated with this notification.
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`

	// The locations relevant to this notification.
	Locations []*Location `json:"locations"`

	// A value specifying the severity level of the notification.
	Level string `json:"level,omitempty"`
}

// NewNotification - creates a new
func NewNotification() *Notification {
	return &Notification{
		Locations: make([]*Location, 0),
	}
}

// WithThreadID - add a ThreadID to the Notification
func (t *Notification) WithThreadID(threadId int) *Notification {
	t.ThreadID = threadId
	return t
}

// WithTimeUtc - add a TimeUtc to the Notification
func (t *Notification) WithTimeUtc(timeUtc string) *Notification {
	t.TimeUtc = timeUtc
	return t
}

// WithException - add a Exception to the Notification
func (e *Notification) WithException(exception *Exception) *Notification {
	e.Exception = exception
	return e
}

// WithProperties - add a Properties to the Notification
func (p *Notification) WithProperties(properties *PropertyBag) *Notification {
	p.Properties = properties
	return p
}

// WithMessage - add a Message to the Notification
func (m *Notification) WithMessage(message *Message) *Notification {
	m.Message = message
	return m
}

// WithDescriptor - add a Descriptor to the Notification
func (d *Notification) WithDescriptor(descriptor *ReportingDescriptorReference) *Notification {
	d.Descriptor = descriptor
	return d
}

// WithAssociatedRule - add a AssociatedRule to the Notification
func (a *Notification) WithAssociatedRule(associatedRule *ReportingDescriptorReference) *Notification {
	a.AssociatedRule = associatedRule
	return a
}

// WithLocations - add a Locations to the Notification
func (l *Notification) WithLocations(locations []*Location) *Notification {
	l.Locations = locations
	return l
}

// AddLocation - add a single Location to the Notification
func (l *Notification) AddLocation(location *Location) *Notification {
	l.Locations = append(l.Locations, location)
	return l
}

// WithLevel - add a Level to the Notification
func (l *Notification) WithLevel(level string) *Notification {
	l.Level = level
	return l
}
