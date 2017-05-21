package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/http"
	)

var D = 500
var workers = 60

func generate_num() int {
	return rand.Intn(3)
}

func bourre() {
	urls := []string{"http://localhost:8080/", "http://localhost:8080/page1/", "http://localhost:8080/x/", "http://localhost:8080/page1/x/"}
	num := generate_num()
	if ( num > 0 ) {
		randurl := rand.Intn(len(urls))
		url := urls[randurl]
		go http.Get(url)
	    go fmt.Println(url)
	}
}

func main() {
	duration := time.Duration(D)*time.Millisecond
	nw := 0
	for _ = range time.Tick(duration) {
		nw = 0
		for nw < workers {
			go bourre()
			nw = nw+1
		}
	}
}