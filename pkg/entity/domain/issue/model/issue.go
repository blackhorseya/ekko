package model

import (
	"encoding/json"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (x *Ticket) MarshalJSON() ([]byte, error) {
	type Alias Ticket

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

func (x *Ticket) UnmarshalJSON(bytes []byte) error {
	type Alias Ticket

	var v struct {
		*Alias
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	id, err := strconv.Atoi(v.ID)
	if err != nil {
		return err
	}

	createdAt, err := time.Parse(time.RFC3339, v.CreatedAt)
	if err != nil {
		return err
	}

	updatedAt, err := time.Parse(time.RFC3339, v.UpdatedAt)
	if err != nil {
		return err
	}

	*x = Ticket{
		Id:        int64(id),
		Title:     v.Title,
		Status:    v.Status,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}

	return nil
}
