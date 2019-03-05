package main

import (
	"net/http"
	n "netdemo"
)

func main() {
	http.HandleFunc("/", n.HIndex)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
