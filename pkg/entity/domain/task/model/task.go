package model

import (
	"encoding/json"
	"strconv"
)

func (x *Task) MarshalJSON() ([]byte, error) {
	type Alias Task

	return json.Marshal(&struct {
		*Alias
		Id string `json:"id"`
	}{
		Alias: (*Alias)(x),
		Id:    strconv.Itoa(int(x.Id)),
	})
}
