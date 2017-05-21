package rethinkdb

import (
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
	return nil
}

func Save(hits []*types.Hit) (int, *terr.Trace) {
	session := conn
	num := 0
	for _, hit := range hits {
		_, err := r.DB("metrics").Table("hits").Insert(hit, r.InsertOpts{Durability: "soft", ReturnChanges: false}).Run(session)
		if err != nil {
			tr := terr.New("db.rethinkdb.Save", err)
			return num, tr
		}
		num++
	}
	return num, nil
}

func connect() (*r.Session, *terr.Trace) {
	user := state.Conf.User
	pwd := state.Conf.Pwd
	addr := state.Conf.Addr
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
