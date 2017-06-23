package state

import (
	"github.com/synw/hitsmon/conf"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
)

var Conf *types.Conf
var Verbosity int

func InitState(dev bool, verbosity int, config ...*types.Conf) *terr.Trace {
	Verbosity = verbosity
	// config
	if len(config) == 1 {
		Conf = config[0]
		return nil
	}
	cf, tr := conf.GetConf(dev, verbosity)
	if tr != nil {
		return tr
	}
	conf.Verify(cf)
	Conf = cf
	return nil
}
