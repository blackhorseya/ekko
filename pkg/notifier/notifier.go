//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package notifier

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// Notifier is an interface for sending notifications.
type Notifier interface {
	SendText(ctx contextx.Contextx, message string) error
}
