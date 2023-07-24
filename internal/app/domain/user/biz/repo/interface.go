//go:generate mockgen -destination=./mock_${GOFILE} -package=repo -source=${GOFILE}

package repo

import (
	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/google/wire"
)

// IRepo declare user repo interface
type IRepo interface {
	// GetProfileByUsername serve caller to given username to get profile from users table
	GetProfileByUsername(ctx contextx.Contextx, username string) (info *userM.Profile, err error)

	// GetProfileByID serve caller to given id to get profile from users table
	GetProfileByID(ctx contextx.Contextx, id int64) (info *userM.Profile, err error)

	// Register serve caller to given profile to insert into users table
	Register(ctx contextx.Contextx, who *userM.Profile) (info *userM.Profile, err error)

	// UpdateToken serve caller to given user and token to update token into users table
	UpdateToken(ctx contextx.Contextx, who *userM.Profile, token string) (info *userM.Profile, err error)
}

// MariadbSet is a provider set for mariadb implementation
var MariadbSet = wire.NewSet(NewMariadb)
