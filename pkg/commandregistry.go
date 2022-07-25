package pkg

import (
	"github.com/trypophob1a/fileschecker/pkg/commands/check"
	"github.com/trypophob1a/fileschecker/pkg/commands/filesnamecollector"
	"github.com/trypophob1a/fileschecker/pkg/commands/selfcheck"
	"github.com/trypophob1a/fileschecker/pkg/core"
	"github.com/trypophob1a/fileschecker/pkg/strategy/checkfinder"
)

func CommandRegistry(r *core.CommandRecorder) *core.CommandRecorder {
	r.AddExecutor("collect", filesnamecollector.NewCollector())
	r.AddExecutor("selfcheck", selfcheck.NewSelfCheck())
	r.AddExecutor("check", check.NewCheck().SetFinder(checkfinder.NewHashmapFinder()))

	return r
}
