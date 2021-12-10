package sarif

type Exception struct {
	PropertyBag
	InnerExceptions []*Exception `json:"innerExceptions,omitempty"`
	Kind            *string      `json:"kind,omitempty"`
	Message         *string      `json:"message,omitempty"`
	Stack           *Stack       `json:"stack,omitempty"`
}

func NewException() *Exception {
	return &Exception{}
}

func (exception *Exception) WithMessage(message string) *Exception {
	exception.Message = &message
	return exception
}

func (exception *Exception) WithKind(kind string) *Exception {
	exception.Kind = &kind
	return exception
}

func (exception *Exception) WithStack(stack Stack) *Exception {
	exception.Stack = &stack
	return exception
}

func (exception *Exception) WithInnerExceptions(exceptions []*Exception) *Exception {
	exception.InnerExceptions = exceptions
	return exception
}

func (exception *Exception) AddInnerException(toAdd *Exception) {
	exception.InnerExceptions = append(exception.InnerExceptions, toAdd)
}
