package pkg

import (
	"github.com/trypophob1a/fileschecker/pkg/commands/check"
	"github.com/trypophob1a/fileschecker/pkg/commands/filesnamecollector"
	"github.com/trypophob1a/fileschecker/pkg/core"
	"github.com/trypophob1a/fileschecker/pkg/strategy"
)

func CommandRegistry(r *core.CommandRecorder) *core.CommandRecorder {
	r.AddExecutor(filesnamecollector.NewCollector())
	r.AddExecutor(check.NewCheck().SetFinder(strategy.NewHashmapFinder()))

	return r
}
