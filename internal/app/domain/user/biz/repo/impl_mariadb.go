package repo

import (
	"database/sql"
	"errors"
	"time"

	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo/dao"
	"github.com/blackhorseya/ekko/pkg/contextx"
	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb serve caller to create an IRepo
func NewMariadb(rw *sqlx.DB) IRepo {
	return &mariadb{
		rw: rw,
	}
}

func (m *mariadb) GetProfileByUsername(ctx contextx.Contextx, username string) (info *um.Profile, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()

	stmt := `select id, username, password, token, created_at, updated_at from users where username = ?`

	var ret dao.Profile
	err = m.rw.GetContext(timeout, &ret, stmt, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return ret.ToEntity(), nil
}

func (m *mariadb) GetProfileByID(ctx contextx.Contextx, id int64) (info *um.Profile, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()

	stmt := `select id, username, password, token, created_at, updated_at from users where id = ?`

	var ret dao.Profile
	err = m.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return ret.ToEntity(), nil
}

func (m *mariadb) Register(ctx contextx.Contextx, who *um.Profile) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()

	stmt := `insert into users (id, username, password, token, created_at, updated_at) values (:id, :username, :password, :token, :created_at, :updated_at)`

	who.CreatedAt = timestamppb.Now()
	who.UpdatedAt = timestamppb.Now()
	_, err := m.rw.NamedExecContext(timeout, stmt, dao.NewProfile(who))
	if err != nil {
		return err
	}

	return nil
}