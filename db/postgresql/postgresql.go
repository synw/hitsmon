package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/synw/hitsmon/state"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
)

var db *gorm.DB

func InitDb() *terr.Trace {
	cn, tr := connect()
	if tr != nil {
		tr := terr.Pass("db.postgresql.InitDb", tr)
		return tr
	}
	db = cn
	return nil
}

func connect() (*gorm.DB, *terr.Trace) {
	db, err := gorm.Open("postgres", "host="+state.Conf.Addr+" user="+state.Conf.User+" dbname="+state.Conf.Db+" sslmode=disable password="+state.Conf.Pwd+"")
	if err != nil {
		tr := terr.New("db.postgresql.connect()", err)
		return db, tr
	}
	return db, nil
}

func Save(hits []*types.Hit) (int, *terr.Trace) {
	num := 0
	for _, hit := range hits {
		db.NewRecord(hit)
		db.Create(hit)
		num++
	}
	return num, nil
}
