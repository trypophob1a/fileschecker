package core

import (
	"github.com/trypophob1a/fileschecker/pkg/interfaces"
)

type CommandRecorder struct {
	commands map[string]func()
}

func NewCommandRecorder() *CommandRecorder {
	return &CommandRecorder{commands: make(map[string]func())}
}

func (r *CommandRecorder) Add(command string, action func()) {
	r.commands[command] = action
}

func (r *CommandRecorder) AddExecutor(command string, executor interfaces.Executed) {
	r.Add(command, executor.Execute)
}

func (r *CommandRecorder) Listener(command string) {
	if action, ok := r.commands[command]; ok {
		action()
	}
}

func (r CommandRecorder) getCommands() map[string]func() {
	return r.commands
}
