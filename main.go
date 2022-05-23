/*
This file is part of vdcsspp
Copyright 2022 epiccakeking
Licensed under 0BSD
*/

/*
Program to generate a CSS theme for AoPS.
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	queue := make(chan []byte, 100)
	go func() {
		for {
			queue <- []byte(Generate())
		}
	}()
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			fmt.Fprintf(w, "Queued: %d", len(queue))
		case "/random.css":
			w.Header().Set("Content-Type", "text/css")
			w.Write(<-queue)
		default:
			http.NotFound(w, r)
		}
	}))
}
