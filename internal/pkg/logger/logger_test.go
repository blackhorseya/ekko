package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSetLevel(t *testing.T) {
	type args struct {
		lvl string
	}
	tests := []struct {
		name      string
		args      args
		wantLevel string
	}{
		{
			name:      "test then info",
			args:      args{"test"},
			wantLevel: "info",
		},
		{
			name:      "debug then debug",
			args:      args{"debug"},
			wantLevel: "debug",
		},
		{
			name:      "\"\" then info",
			args:      args{""},
			wantLevel: "info",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.lvl)
			got := logrus.GetLevel().String()
			if got != tt.wantLevel {
				t.Errorf("SetLevel() = %v, want %v", got, tt.wantLevel)
			}
		})
	}
}
