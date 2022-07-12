package core

import (
	"reflect"
	"strings"

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

func (r *CommandRecorder) AddExecutor(command interfaces.Executed) {
	r.Add(r.getCommandName(reflect.TypeOf(command).String()), command.Execute)
}

func (r *CommandRecorder) getCommandName(command string) string {
	for i, r := range command {
		if r == '.' {
			return strings.ToLower(command[i+8:])
		}
	}
	return command
}

func (r *CommandRecorder) Listener(command string) {
	if action, ok := r.commands[command]; ok {
		action()
	}
}

func (r CommandRecorder) getCommands() map[string]func() {
	return r.commands
}
