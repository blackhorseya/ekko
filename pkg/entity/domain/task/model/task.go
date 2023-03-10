package model

import (
	"encoding/json"
	"strconv"
	"time"
)

func (x *Task) MarshalJSON() ([]byte, error) {
	type Alias Task

	return json.Marshal(&struct {
		*Alias
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias:     (*Alias)(x),
		ID:        strconv.Itoa(int(x.Id)),
		CreatedAt: x.CreatedAt.AsTime().UTC().Format(time.RFC3339),
		UpdatedAt: x.UpdatedAt.AsTime().UTC().Format(time.RFC3339),
	})
}
