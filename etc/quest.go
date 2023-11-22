package etc

import (
	"time"
)

type Quest struct {
	Description string
	Conditions  func() bool
	Created     time.Time
}
