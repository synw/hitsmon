package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
)

var dB *gorm.DB

func InitDb(db *types.Db) *terr.Trace {
	cn, tr := connect(db)
	if tr != nil {
		tr := terr.Pass("db.postgresql.InitDb", tr)
		return tr
	}
	dB = cn
	return nil
}

func connect(cdb *types.Db) (*gorm.DB, *terr.Trace) {
	db, err := gorm.Open("postgres", "host="+cdb.Addr+" user="+cdb.User+" dbname="+cdb.Name+" sslmode=disable password="+cdb.Pwd+"")
	if err != nil {
		tr := terr.New("db.postgresql.connect()", err)
		return db, tr
	}
	return db, nil
}

func Save(hits []*types.Hit) (int, *terr.Trace) {
	num := 0
	for _, hit := range hits {
		dB.NewRecord(hit)
		dB.Create(hit)
		num++
	}
	return num, nil
}
