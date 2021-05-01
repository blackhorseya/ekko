package trace

import (
	"regexp"
	"testing"
)

func TestNewTraceID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "generate then match regex",
			want: "trace-id-\\d+-\\d{4}\\.\\d{2}\\.\\d{2}\\.\\d{2}\\.\\d{2}\\.\\d{2}\\.\\d+-\\d+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTraceID()
			if match, _ := regexp.MatchString(tt.want, got); !match {
				t.Errorf("NewTraceID() = %v, want %v", got, tt.want)
			}
		})
	}
}
