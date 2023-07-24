package testdata

import (
	"time"

	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Profile1 user 1
	Profile1 = &userM.Profile{
		Id:        1,
		Username:  "username",
		Password:  "password",
		Token:     "token",
		CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}
)
