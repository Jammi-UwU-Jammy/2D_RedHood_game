package etc

import (
	"time"
)

type Quest struct {
	Title       string
	Description string
	Conditions  func() bool
	Created     time.Time
}
