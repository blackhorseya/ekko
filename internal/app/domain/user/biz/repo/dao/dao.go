package dao

import (
	"time"

	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Profile serve caller to provide profile dao
type Profile struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewProfile(info *userM.Profile) *Profile {
	return &Profile{
		ID:        info.Id,
		Username:  info.Username,
		Password:  info.Password,
		Token:     info.Token,
		CreatedAt: info.CreatedAt.AsTime().UTC(),
		UpdatedAt: info.UpdatedAt.AsTime().UTC(),
	}
}

// ToEntity serve caller to convert to Profile entity
func (p *Profile) ToEntity() *userM.Profile {
	return &userM.Profile{
		Id:        p.ID,
		Username:  p.Username,
		Password:  p.Password,
		Token:     p.Token,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}
