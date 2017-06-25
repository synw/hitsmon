package influxdb

import (
	"github.com/influxdata/influxdb/client/v2"
	//"github.com/synw/hitsmon/db/dbutils"
	"github.com/synw/hitsmon/types"
	"github.com/synw/terr"
	"time"
)

var cli client.Client

func Init(db *types.Db) *terr.Trace {
	var err error
	terr.Debug(db)

	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     db.Addr,
		Username: db.User,
		Password: db.Pwd,
	})
	if err != nil {
		tr := terr.New("db.influxdb.Save", err)
		return tr
	}
	return nil
}

func Save(db *types.Db, hits []*types.Hit) (int, *terr.Trace) {
	num := 0
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db.Name,
		Precision: "ms",
	})
	if err != nil {
		tr := terr.New("db.influxdb.Save", err)
		return num, tr
	}
	// Create a point and add to batch
	tags := make(map[string]string)
	/*if aggregate {
		domain := hits[0].Domain
		tags = map[string]string{"service": "hitsmon", "domain": domain}
		all, users, anonymous := dbutils.Aggregate(hits)
		fields := map[string]interface{}{
			"all":       all,
			"users":     users,
			"anonymous": anonymous,
		}
		pt, err := client.NewPoint("hits", tags, fields, time.Now())
		if err != nil {
			tr := terr.New("db.influxdb.Save", err)
			return 0, tr
		}
		bp.AddPoint(pt)
		num = all
	} else {*/
	for _, hit := range hits {
		tags = map[string]string{
			"service":       "hitsmon",
			"domain":        hit.Domain,
			"user":          hit.User,
			"path":          hit.Path,
			"referer":       hit.Referer,
			"user_agent":    hit.UserAgent,
			"method":        hit.Method,
			"authenticated": hit.IsAuthenticated,
			"staff":         hit.IsStaff,
			"superuser":     hit.IsSuperuser,
			"status_code":   hit.StatusCode,
			"view":          hit.View,
			"module":        hit.Module,
			"ip":            hit.Ip,
		}
		fields := map[string]interface{}{
			"num":            1,
			"request_time":   hit.RequestTime,
			"content_length": hit.ContentLength,
			"num_queries":    hit.NumQueries,
			"queries_time":   hit.QueriesTime,
		}
		pt, err := client.NewPoint("hits", tags, fields, time.Now())
		if err != nil {
			tr := terr.New("db.influxdb.Save", err)
			return num, tr
		}
		bp.AddPoint(pt)
		num++
	}
	//}
	// Write the batch
	if err := cli.Write(bp); err != nil {
		tr := terr.New("db.influxdb.Save", err)
		return 0, tr
	}
	return num, nil
}
