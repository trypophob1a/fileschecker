package testdata

type commandExecutor struct {
	name string
}

func NewExecutor(name string) *commandExecutor {
	return &commandExecutor{name: name}
}

func (e commandExecutor) Execute() {
	println(e.name)
}
