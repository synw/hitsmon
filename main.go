package main

import (
	"flag"
	"fmt"
	"github.com/synw/hitsmon/db"
	"github.com/synw/hitsmon/db/redis"
	"github.com/synw/hitsmon/log"
	"github.com/synw/hitsmon/state"
	"time"
)

var dev = flag.Bool("d", false, "Dev mode")
var verbosity = flag.Int("v", 0, "Verbosity")

func main() {
	flag.Parse()
	// init state
	tr := state.InitState(*dev, *verbosity)
	if tr != nil {
		tr.Fatal("initializing config")
	}
	// init db
	tr = db.Init(*verbosity, state.Conf.Db)
	if tr != nil {
		tr.Fatal("initializing database")
	}
	if state.Verbosity > 0 {
		fmt.Println("Starting to work ...")
	}
	// init logger
	tr = log.Init(state.Conf)
	if tr != nil {
		tr.Printf("main")
	}
	// now work
	if *verbosity > 0 {
		fmt.Println("Starting hitsmon service with database " + state.Conf.Db.Name + " ...")
	}
	for {
		duration := time.Duration(state.Conf.Frequency) * time.Second
		for range time.Tick(duration) {
			go redis.ProcessHits()
		}
	}
}
