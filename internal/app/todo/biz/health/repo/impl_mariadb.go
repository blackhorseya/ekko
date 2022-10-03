package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/jmoiron/sqlx"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb is a constructor health of implement repo
func NewMariadb(rw *sqlx.DB) IHealthRepo {
	return &mariadb{rw: rw}
}

// Ping sends a ping command to verify that the client can connect to the deployment
func (i *mariadb) Ping(ctx contextx.Contextx, timeout time.Duration) error {
	withTimeout, cancel := contextx.WithTimeout(ctx, timeout)
	defer cancel()

	err := i.rw.PingContext(withTimeout)
	if err != nil {
		return err
	}

	return nil
}
