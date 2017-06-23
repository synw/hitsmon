package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/synw/hitsmon/db"
	"github.com/synw/hitsmon/state"
	"log"
	"strconv"
	"time"
)

var Conn = connect()

func connect() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func ProcessHits() error {
	// get hits set
	prefix := state.Conf.Domain + "_hit*"
	keys, err := redis.Values(Conn.Do("KEYS", prefix))
	if err != nil {
		fmt.Println("KEYS: error retrieving Redis keys:", err)
	}
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	now := time.Now()
	date := strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute()) + ":" + strconv.Itoa(now.Second())
	if len(keys) > 0 {
		values, err := redis.Strings(Conn.Do("MGET", args...))
		if err != nil {
			fmt.Println("MGET: error retrieving Redis values:", err)
		}
		// save the keys into the db
		go db.Save(values)
		// delete the recorded keys from Redis
		Conn.Send("MULTI")
		for i, _ := range keys {
			//fmt.Println(keys[i])
			Conn.Send("DEL", keys[i])
		}
		res, err := Conn.Do("EXEC")
		if err != nil {
			if state.Verbosity > 0 {
				fmt.Println("DEL: error deleting Redis keys:", err)
			}
		}
		// then report
		if state.Verbosity > 0 {
			fmt.Println(date, "-", len(res.([]interface{})[:]), "hits")
		}
	} else {
		if state.Verbosity > 0 {
			fmt.Println(date, "- 0 hits")
		}
	}
	return nil
}
