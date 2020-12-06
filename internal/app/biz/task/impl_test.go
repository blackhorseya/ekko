package task

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/entities"
)

func Test_impl_Create(t *testing.T) {
	type args struct {
		t *entities.Task
	}
	tests := []struct {
		name     string
		args     args
		wantTask *entities.Task
		wantErr  bool
	}{
		{
			name: "missing title then nil true",
			args: args{&entities.Task{
				Title: "",
			}},
			wantTask: nil,
			wantErr:  true,
		},
		// todo: 2020-12-06|13:34|doggy|test it using testify
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotTask, err := i.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}
