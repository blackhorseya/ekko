//go:build external

package restful

import (
	"net/http"
	"testing"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	"go.uber.org/zap"
)

func TestHealthz(t *testing.T) {
	logger := zap.NewExample()

	cfg := config.NewDefaultConfig()
	service, err := NewService(cfg, logger)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Start()
	if err != nil {
		t.Fatal(err)
	}

	client := http.DefaultClient
	resp, err := client.Get("http://localhost:1992/api/healthz")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("status code not ok")
	}
}
