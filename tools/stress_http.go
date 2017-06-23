package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var D = 500
var workers = 10

func generate_num() int {
	return rand.Intn(3)
}

func bourre() {
	urls := []string{"http://localhost:8000/", "http://localhost:8000/page1/", "http://localhost:8000/contact/", "http://localhost:8000/page2/"}
	num := generate_num()
	if num > 0 {
		randurl := rand.Intn(len(urls))
		url := urls[randurl]
		go http.Get(url)
		go fmt.Println(url)
	}
}

func main() {
	duration := time.Duration(D) * time.Millisecond
	nw := 0
	for _ = range time.Tick(duration) {
		nw = 0
		for nw < workers {
			go bourre()
			nw = nw + 1
		}
	}
}
