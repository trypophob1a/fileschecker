package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/trypophob1a/fileschecker/pkg"
	"github.com/trypophob1a/fileschecker/pkg/config"
	"github.com/trypophob1a/fileschecker/pkg/core"
)

func main() {
	commandName := kingpin.MustParse(config.GetApp().Parse(os.Args[1:]))
	recorder := pkg.CommandRegistry(core.NewCommandRecorder())
	recorder.Listener(commandName)
}
