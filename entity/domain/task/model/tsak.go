package model

import (
	"encoding/json"
	"time"
)

func (x *Ticket) MarshalJSON() ([]byte, error) {
	type Alias Ticket

	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias:     (*Alias)(x),
		CreatedAt: x.CreatedAt.AsTime().UTC().Format(time.RFC3339),
		UpdatedAt: x.UpdatedAt.AsTime().UTC().Format(time.RFC3339),
	})
}
