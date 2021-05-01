package health

import (
	"time"

	repository2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
)

type impl struct {
	HealthRepo repository2.IRepo
}

// NewImpl is a constructor of implement business with parameters
func NewImpl(healthRepo repository2.IRepo) IBiz {
	return &impl{HealthRepo: healthRepo}
}

// Readiness to handle application has been ready
func (i *impl) Readiness() (ok bool, err error) {
	err = i.HealthRepo.Ping(2 * time.Second)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness() (ok bool, err error) {
	err = i.HealthRepo.Ping(5 * time.Second)
	if err != nil {
		return false, err
	}

	return true, nil
}
