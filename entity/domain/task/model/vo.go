package model

import (
	"encoding/json"
	"fmt"
)

// TicketStatus is an interface that represents the status of a ticket.
type TicketStatus interface {
	fmt.Stringer
	json.Marshaler
	json.Unmarshaler
}
