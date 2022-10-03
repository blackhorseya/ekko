package health

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"go.uber.org/zap"
)

type impl struct {
	repo repo.IHealthRepo
}

// NewImpl is a constructor of implement business with parameters
func NewImpl(healthRepo repo.IHealthRepo) IBiz {
	return &impl{
		repo: healthRepo,
	}
}

// Readiness to handle application has been ready
func (i *impl) Readiness(ctx contextx.Contextx) (ok bool, err error) {
	err = i.repo.Ping(ctx, 2*time.Second)
	if err != nil {
		ctx.Error(er.ErrPing.Error(), zap.Error(err))
		return false, er.ErrPing
	}

	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness(ctx contextx.Contextx) (ok bool, err error) {
	err = i.repo.Ping(ctx, 5*time.Second)
	if err != nil {
		ctx.Error(er.ErrPing.Error(), zap.Error(err))
		return false, er.ErrPing
	}

	return true, nil
}
