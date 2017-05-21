package conf

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
)

func GetConf(dev bool, verbosity int) (*types.Conf, *terr.Trace) {
	var conf *types.Conf
	// set some defaults for conf
	if dev {
		viper.SetConfigName("dev_config")
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(".")
	viper.SetDefault("type", "")
	viper.SetDefault("addr", "localhost:28015")
	viper.SetDefault("user", "")
	viper.SetDefault("password", "")
	viper.SetDefault("db", "")
	viper.SetDefault("table", "")
	viper.SetDefault("frequency", 1)
	viper.SetDefault("separator", "#!#")
	// get the actual conf
	err := viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigParseError:
			trace := terr.New("conf.getConf", err)
			return conf, trace
		default:
			err := errors.New("Unable to locate config file")
			trace := terr.New("conf.getConf", err)
			return conf, trace
		}
	}
	dbtype := viper.Get("type").(string)
	addr := viper.Get("addr").(string)
	user := viper.Get("user").(string)
	pwd := viper.Get("password").(string)
	db := viper.Get("db").(string)
	table := viper.Get("table").(string)
	frequency := viper.GetInt("frequency")
	separator := viper.Get("separator").(string)
	if dbtype == "" {
		err := errors.New("Please set the database type into your config file: ex: \"type\":\"rethinkdb\"")
		tr := terr.New("conf.GetConf", err)
		terr.Fatal("loading configuration", tr)
	}
	endconf := &types.Conf{dbtype, addr, user, pwd, db, table, frequency, separator, dev, verbosity}
	verifyConf(endconf)
	return endconf, nil
}

func verifyConf(conf *types.Conf) {
	if conf.DbType == "" {
		err := errors.New("No database type specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
	if conf.Db == "" {
		err := errors.New("No database specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
	if conf.Table == "" {
		err := errors.New("No database table specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
}
