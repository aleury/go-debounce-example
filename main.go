package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	ch := make(chan string)

	go debounce(1*time.Second, ch, func(msg string) {
		log.Println(msg)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		ch <- q.Get("msg")
		w.Write([]byte("ok!\n"))
	})

	http.ListenAndServe(":8000", nil)
}

func debounce(d time.Duration, ch chan string, f func(msg string)) {
	var msg string
	timer := time.NewTimer(d)
	for {
		select {
		case msg = <-ch:
			timer.Reset(d)
			msg = strings.TrimSpace(msg)
		case <-timer.C:
			if msg != "" {
				f(msg)
			}
		}
	}
}
