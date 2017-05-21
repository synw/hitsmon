package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"math/rand"
	"strconv"
	"time"
)

func generate_num() int {
	return rand.Intn(2)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	frequency := 1
	duration := time.Duration(frequency) * time.Millisecond
	i := 0
	for x := range time.Tick(duration) {
		num := generate_num()
		if num > 0 {
			t := time.Now()
			tu := t.UnixNano()
			timestamp_str := strconv.Itoa(int(tu))
			key := "hit_" + timestamp_str
			val := "localhost#!#/contact/rest/#!#GET#!#127.0.0.1#!#Mozilla/5.0 (X11; NetBSD amd64; rv:48.0) Gecko/20100101 Firefox/48.0#!#syn#!#http://127.0.0.1:8000/page2/"
			go client.Set(key, val, 0).Err()
			go fmt.Println("Hit", i, x)
			i = i + 1
		} else {
			go fmt.Println("x")
		}
	}
}
