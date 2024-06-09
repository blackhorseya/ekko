//go:build external

package rest

import (
	"testing"

	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	err := configx.LoadConfig("")
	if err != nil {
		t.Fatalf("configx.LoadConfig() error = %v", err)
	}

	service, err := New(viper.GetViper())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = service.Start()
	if err != nil {
		t.Fatalf("service.Start() error = %v", err)
	}

	err = service.AwaitSignal()
	if err != nil {
		t.Fatalf("service.AwaitSignal() error = %v", err)
	}
}
