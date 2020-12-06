package health

import "testing"

func Test_impl_Readiness(t *testing.T) {
	tests := []struct {
		name    string
		wantOk  bool
		wantErr bool
	}{
		{
			name:    "readiness then true nil",
			wantOk:  true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotOk, err := i.Readiness()
			if (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Readiness() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_impl_Liveness(t *testing.T) {
	tests := []struct {
		name    string
		wantOk  bool
		wantErr bool
	}{
		{
			name:    "liveness then true nil",
			wantOk:  true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotOk, err := i.Liveness()
			if (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Liveness() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}