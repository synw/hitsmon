package main

import (
	"flag"
	"fmt"
	"github.com/synw/hitsmon/db"
	"github.com/synw/hitsmon/db/redis"
	//"github.com/synw/hitsmon/log"
	"github.com/synw/hitsmon/state"
	"sync"
	"time"
)

var dev = flag.Bool("d", false, "Dev mode")
var verbosity = flag.Int("v", 0, "Verbosity")

func main() {
	flag.Parse()
	// init state
	conf, tr := state.InitState(*dev, *verbosity)
	if tr != nil {
		tr.Fatal("initializing config")
	}
	//fmt.Println("\n-----------", conf, "\n-----------")
	// init db
	tr = db.Init(conf.Db, conf.Separator, *verbosity)
	if tr != nil {
		tr.Fatal("initializing database")
	}
	if *verbosity > 0 {
		fmt.Println("Starting to work ...")
	}
	// init logger
	/*tr = log.Init(state.Conf)
	if tr != nil {
		tr.Printf("main")
	}*/
	// now work
	if *verbosity > 0 {
		//fmt.Println("Starting hitsmon service with database " + conf.Db.Name + " ...")
	}
	var mutex = &sync.Mutex{}
	for {
		duration := time.Duration(conf.Frequency) * time.Second
		for range time.Tick(duration) {
			go redis.ProcessHits(conf.Domain, mutex, *verbosity)
		}
	}
}
