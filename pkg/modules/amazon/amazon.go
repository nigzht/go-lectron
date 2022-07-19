package amazon

import (
	"fmt"

	"github.com/nemesis-io/tasks"
)

type AmazonInput struct{}

type Amazon struct{}

func NewAmazon() *Amazon {
	return &Amazon{}
}

func (a *Amazon) NewTask(i AmazonInput) tasks.Task {
	panic(fmt.Errorf("unimplemented"))
}
