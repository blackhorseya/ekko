package task

import (
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/task/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	taskBiz     *mocks.Biz
	taskHandler IHandler
}

func (s *handlerTestSuite) SetupTest() {
	s.taskBiz = new(mocks.Biz)
	handler, err := CreateTaskHandler(s.taskBiz)
	if err != nil {
		panic(err)
	}
	s.taskHandler = handler
}

func (s *handlerTestSuite) TearDownTest() {
	s.taskBiz.AssertExpectations(s.T())
}

func (s *handlerTestSuite) Test_impl_List() {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		mockFunc func()
	}{
		{
			name: "",
			args: args{},
		},
	}
	for _, _ = range tests {
	}
}

func TestTaskHandler(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
