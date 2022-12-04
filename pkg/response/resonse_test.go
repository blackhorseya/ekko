package response

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteTester struct {
	suite.Suite
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) TestResponse_WithMessage() {
	type fields struct {
		Code int
		Msg  string
		Data interface{}
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		{
			name:   "append message then success",
			fields: fields{Code: 200, Data: nil},
			args:   args{message: "message"},
			want: &Response{
				Code: 200,
				Msg:  "message",
				Data: nil,
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Code: tt.fields.Code,
				Msg:  tt.fields.Msg,
				Data: tt.fields.Data,
			}
			if got := resp.WithMessage(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *suiteTester) TestResponse_WithData() {
	type fields struct {
		Code int
		Msg  string
		Data interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		{
			name:   "append data then success",
			fields: fields{Code: 200, Msg: "message", Data: nil},
			args:   args{data: "data"},
			want:   &Response{Code: 200, Msg: "message", Data: "data"},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Code: tt.fields.Code,
				Msg:  tt.fields.Msg,
				Data: tt.fields.Data,
			}
			if got := resp.WithData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithData() = %v, want %v", got, tt.want)
			}
		})
	}
}
