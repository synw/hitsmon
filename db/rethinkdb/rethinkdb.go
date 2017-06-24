package rethinkdb

import (
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	r "gopkg.in/dancannon/gorethink.v3"
)

var conn *r.Session
var dB *types.Db

func InitDb(db *types.Db) *terr.Trace {
	cn, tr := connect(db)
	if tr != nil {
		tr := terr.Pass("db.rethinkdb.InitDb", tr)
		return tr
	}
	conn = cn
	dB = db
	return nil
}

func Save(hits []*types.Hit) (int, *terr.Trace) {
	session := conn
	num := 0
	for _, hit := range hits {
		_, err := r.DB(dB.Name).Table(dB.Table).Insert(hit, r.InsertOpts{Durability: "soft", ReturnChanges: false}).Run(session)
		if err != nil {
			tr := terr.New("db.rethinkdb.Save", err)
			return num, tr
		}
		num++
	}
	return num, nil
}

func connect(db *types.Db) (*r.Session, *terr.Trace) {
	user := db.User
	pwd := db.Pwd
	addr := db.Addr
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
