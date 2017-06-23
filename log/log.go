package log

import (
	"github.com/Abramovic/logrus_influxdb"
	"github.com/Sirupsen/logrus"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	"io/ioutil"
	"time"
)

var logger = logrus.New()

func Init(conf *types.Conf) *terr.Trace {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	config := &logrus_influxdb.Config{
		Host:          conf.Db.Host,
		Port:          conf.Db.Port,
		Database:      conf.Db.Name,
		UseHTTPS:      false,
		Precision:     "ns",
		Tags:          []string{"hits"},
		BatchInterval: (5 * time.Second),
		BatchCount:    0, // set to "0" to disable batching
	}
	hook, err := logrus_influxdb.NewInfluxDB(config)
	if err == nil {
		logger.Hooks.Add(hook)
	} else {
		tr := terr.New("log.Init", err)
		return tr
	}
	logger.Out = ioutil.Discard
	return nil
}

func NewFromHitsPack(hits []*types.Hit, anonymous_hits int, user_hits int) {
	level := "info"
	domain := hits[0].Domain
	num := len(hits)
	now := time.Now().UnixNano()
	logobj := logger.WithFields(logrus.Fields{
		"domain":    domain,
		"date":      now,
		"num":       num,
		"anonymous": anonymous_hits,
		"users":     user_hits,
	})
	if level == "debug" {
		logobj.Debug()
	} else if level == "warn" {
		logobj.Warn()
	} else if level == "error" {
		logobj.Error()
	} else if level == "fatal" {
		logobj.Fatal()
	} else if level == "panic" {
		logobj.Panic()
	} else {
		logobj.Info()
	}
}
