package sarif

import "github.com/google/uuid"

// Guid - A stable, unique identifier for many entities in the form of a GUID.
type Guid string

func NewGuid() Guid {
	return Guid(uuid.New().String())
}
