package main

import (
	"flag"
	"fmt"
	"github.com/synw/hitsmon/db"
	"github.com/synw/hitsmon/db/redis"
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
	tr = db.Init()
	if tr != nil {
		tr.Fatal("initializing database")
	}
	if state.Verbosity > 0 {
		fmt.Println("Starting to work ...")
	}
	// now work
	for {
		duration := time.Duration(state.Conf.Frequency) * time.Second
		for range time.Tick(duration) {
			go redis.ProcessHits()
		}
	}
}
