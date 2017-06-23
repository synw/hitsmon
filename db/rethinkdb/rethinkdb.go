package rethinkdb

import (
	"fmt"
	"github.com/synw/hitsmon/state"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	r "gopkg.in/dancannon/gorethink.v3"
)

var conn *r.Session

func InitDb() *terr.Trace {
	cn, tr := connect()
	if tr != nil {
		tr := terr.Pass("db.rethinkdb.InitDb", tr)
		return tr
	}
	conn = cn
	if state.Verbosity > 0 {
		fmt.Println("Rethinkdb database is up at ", state.Conf.Db.Addr)
	}
	return nil
}

func Save(hits []*types.Hit) (int, *terr.Trace) {
	session := conn
	num := 0
	for _, hit := range hits {
		_, err := r.DB(state.Conf.Db.Name).Table(state.Conf.Db.Table).Insert(hit, r.InsertOpts{Durability: "soft", ReturnChanges: false}).Run(session)
		if err != nil {
			tr := terr.New("db.rethinkdb.Save", err)
			return num, tr
		}
		num++
	}
	return num, nil
}

func connect() (*r.Session, *terr.Trace) {
	user := state.Conf.Db.User
	pwd := state.Conf.Db.Pwd
	addr := state.Conf.Db.Addr
	// connect to Rethinkdb
	session, err := r.Connect(r.ConnectOpts{
		Address:    addr,
		Username:   user,
		Password:   pwd,
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		tr := terr.New("db.rethinkdb.connect()", err)
		return session, tr
	}
	return session, nil
}
