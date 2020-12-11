package health

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/health/repository"
)

type impl struct {
	HealthRepo repository.HealthRepo
}

// NewImpl is a constructor of implement business with parameters
func NewImpl(healthRepo repository.HealthRepo) Biz {
	return &impl{HealthRepo: healthRepo}
}

// Readiness to handle application has been ready
func (i *impl) Readiness() (ok bool, err error) {
	// todo: 2020-12-11|19:08|doggy|implement me
	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness() (ok bool, err error) {
	// todo: 2020-12-10|10:17|doggy|implement me
	return true, nil
}
