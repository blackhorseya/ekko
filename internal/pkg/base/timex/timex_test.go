package timex

import (
	"reflect"
	"testing"
	"time"
)

var (
	unixNano = int64(1610548520788105000)
)

func TestUnix(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "1610548520788105000 then time",
			args: args{t: unixNano},
			want: time.Unix(unixNano/1e9, unixNano%1e9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unix(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unix() = %v, want %v", got, tt.want)
			}
		})
	}
}
