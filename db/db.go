package db

import (
	"errors"
	"fmt"
	"github.com/SKAhack/go-shortid"
	"github.com/synw/hitsmon/db/influxdb"
	pg "github.com/synw/hitsmon/db/postgresql"
	"github.com/synw/hitsmon/db/rethinkdb"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	"strings"
	"time"
)

var dB *types.Db
var g = shortid.Generator()
var verbosity = 0
var separator string

func Init(db *types.Db, sep string, verb int) *terr.Trace {
	var tr *terr.Trace
	verbosity = verb
	separator = sep
	if db.Type == "rethinkdb" {
		tr = rethinkdb.InitDb(db)
	} else if db.Type == "postgresql" {
		tr = pg.InitDb(db)
	} else if db.Type == "influxdb" {
		tr = influxdb.Init(db)
	} else {
		err := errors.New("No database configured")
		tr = terr.New("db.InitDb", err)
	}
	if tr != nil {
		tr = terr.Pass("db.Init", tr)
		if verbosity > 0 {
			tr.Printf("db.Init")
		}
		return tr
	}
	dB = db
	if verbosity > 0 {
		msg := db.Type + " database is up at " + db.Addr
		terr.Ok(msg)
	}
	return nil
}

func Save(values []string) {
	var hits []*types.Hit
	user_hits := 0
	anonymous_hits := 0
	for _, doc := range values {
		// unpack the data
		data := strings.Split(doc, separator)
		id := g.Generate()
		hit := &types.Hit{
			id,
			data[0],
			data[1],
			data[2],
			data[3],
			data[4],
			data[5],
			data[6],
			data[7],
			data[8],
			data[9],
			data[10],
			data[11],
			data[12],
			data[13],
			data[14],
			data[15],
			data[16],
			data[17],
		}
		hits = append(hits, hit)
		if hit.User == "anonymous" {
			anonymous_hits++
		} else {
			user_hits++
		}
	}
	t1 := time.Now()
	var tr *terr.Trace
	num := 0
	if dB.Type == "rethinkdb" {
		num, tr = rethinkdb.Save(hits)
	} else if dB.Type == "postgresql" {
		num, tr = pg.Save(hits)
	} else if dB.Type == "influxdb" {
		num, tr = influxdb.Save(dB, hits)
	}
	t2 := time.Since(t1)
	if tr != nil {
		if verbosity > 0 {
			tr.Printf("db.Save")
		}
	} else {
		if verbosity > 1 {
			fmt.Println("Saved ", num, "hits in the database in", t2)
		}
	}
}
