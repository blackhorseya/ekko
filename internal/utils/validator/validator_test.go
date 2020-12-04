package validator

import "testing"

func TestContainInt(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,4] 2 then true",
			args: args{[]int{1, 2, 3, 4}, 2},
			want: true,
		},
		{
			name: "[1,2,3,4] 5 then false",
			args: args{[]int{1, 2, 3, 4}, 5},
			want: false,
		},
		{
			name: "[] 5 then false",
			args: args{[]int{}, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("ContainInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
