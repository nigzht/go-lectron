package server

import (
	"context"
	"fmt"
)

type TasksServerImpl struct{}

func NewTasksServerImpl() *TasksServerImpl {
	return &TasksServerImpl{}
}

func (a *TasksServerImpl) AddTask(context.Context, *TaskInput) (*TaskDetails, error) {
	panic(fmt.Errorf("unimplemented"))
}

func (a *TasksServerImpl) mustEmbedUnimplementedTasksServer() {}
