package modules

import (
	"github.com/nemesis-io/tasks"
)

type Module interface {
	Name() string
	NewTask(input any) (tasks.Task, error)
}
