package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"goravel-tinker/view"
)

type Tinker struct{}

func NewTinker() *Tinker {
	return &Tinker{}
}

// Signature The name and signature of the console command.
func (receiver *Tinker) Signature() string {
	return "tinker"
}

// Description The console command description.
func (receiver *Tinker) Description() string {
	return "a tinker command tool"
}

// Extend The console command extend.
func (receiver *Tinker) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Tinker) Handle(ctx console.Context) error {
	view.Call()
	return nil
}
