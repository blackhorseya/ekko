//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	userM "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
)

// IBiz declare user domain interface
type IBiz interface {
	// Signup serve caller to given username and password to create a user
	Signup(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error)

	// Login serve caller to given username and password to login system
	Login(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error)

	// Logout serve caller to given user's profile to logout system
	Logout(ctx contextx.Contextx, who *userM.Profile) error

	// WhoAmI serve caller to given token to get user's profile
	WhoAmI(ctx contextx.Contextx, token string) (info *userM.Profile, err error)
}
