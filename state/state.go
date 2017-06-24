package state

import (
	conflib "github.com/synw/hitsmon/conf"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
)

func InitState(dev bool, verbosity int, config ...*types.Conf) (*types.Conf, *terr.Trace) {
	var conf *types.Conf
	// config
	if len(config) == 1 {
		conf := config[0]
		return conf, nil
	}
	conf, tr := conflib.GetConf(dev, verbosity)
	if tr != nil {
		return conf, tr
	}
	conflib.Verify(conf)
	return conf, nil
}
