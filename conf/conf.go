package conf

import (
	"errors"
	"fmt"
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
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 8086)
	viper.SetDefault("user", "")
	viper.SetDefault("password", "")
	viper.SetDefault("db", "")
	viper.SetDefault("table", "")
	viper.SetDefault("domain", "localhost")
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
	host := viper.Get("host").(string)
	domain := viper.Get("domain").(string)
	port := viper.GetInt("port")
	frequency := viper.GetInt("frequency")
	separator := viper.Get("separator").(string)
	if dbtype == "" {
		err := errors.New("Please set the database type into your config file: ex: \"type\":\"rethinkdb\"")
		tr := terr.New("conf.GetConf", err)
		terr.Fatal("loading configuration", tr)
	}
	database := &types.Db{dbtype, addr, host, port, user, pwd, db, table}
	endconf := &types.Conf{database, frequency, domain, separator, dev, verbosity}
	if verbosity > 2 {
		fmt.Println("Config:\n", endconf)
	}
	return endconf, nil
}

func Verify(conf *types.Conf) {
	if conf.Db.Type == "" {
		err := errors.New("No database type specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
	if conf.Db.Name == "" {
		err := errors.New("No database specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
	if conf.Db.Table == "" {
		err := errors.New("No database table specified")
		tr := terr.New("conf.verifyConf", err)
		terr.Fatal("verifying config", tr)
	}
}
