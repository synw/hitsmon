package db

import (
	"errors"
	"fmt"
	pg "github.com/synw/hitsmon/db/postgresql"
	"github.com/synw/hitsmon/db/rethinkdb"
	"github.com/synw/hitsmon/state"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	"strings"
	"time"
)

func Init() *terr.Trace {
	if state.Conf.DbType == "rethinkdb" {
		tr := rethinkdb.InitDb()
		if tr != nil {
			tr = terr.Pass("db.Init", tr)
			if state.Verbosity > 0 {
				tr.Printf("db.Init")
			}
			return tr
		}
	} else if state.Conf.DbType == "postgresql" {
		tr := pg.InitDb()
		if tr != nil {
			tr = terr.Pass("db.Init", tr)
			if state.Verbosity > 0 {
				tr.Printf("db.Init")
			}
			return tr
		}
	} else {
		err := errors.New("No database configured")
		tr := terr.New("db.InitDb", err)
		return tr
	}
	return nil
}

func Save(values []string) {
	var hits []*types.Hit
	for _, doc := range values {
		//fmt.Println("Doc", doc)
		// unpack the data
		data := strings.Split(doc, state.Conf.Separator)
		datenow := time.Now()
		hit := &types.Hit{0, data[0], data[1], data[2], data[3], data[4], data[5], data[6], datenow}
		hits = append(hits, hit)
	}
	if state.Conf.DbType == "rethinkdb" {
		t1 := time.Now()
		num, tr := rethinkdb.Save(hits)
		t2 := time.Since(t1)
		if tr != nil {
			if state.Verbosity > 0 {
				tr.Printf("db.Save")
			}
		} else {
			if state.Verbosity > 1 {
				fmt.Println("Saved ", num, "hits in the database in", t2)
			}
		}
	} else if state.Conf.DbType == "postgresql" {
		t1 := time.Now()
		num, tr := pg.Save(hits)
		t2 := time.Since(t1)
		if tr != nil {
			if state.Verbosity > 0 {
				tr.Printf("db.Save")
			}
		} else {
			if state.Verbosity > 1 {
				fmt.Println("Saved ", num, "hits in the database in", t2)
			}
		}
	}
}
