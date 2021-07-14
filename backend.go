package main

import (
        "log"
        "net/http"
        "time"
)

var slowdown time.Duration

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                time.Sleep(slowdown)
                w.Write([]byte(`{ "text": "Hello world!"}`))
        })
        http.HandleFunc("/slowdown", func(w http.ResponseWriter, r *http.Request) {
                slowdown = 10000 * time.Millisecond
                w.Write([]byte(`{ "text": "Slowed down server; it now takes 10s to respond"}`))
        })
        http.HandleFunc("/speedup", func(w http.ResponseWriter, r *http.Request) {
                slowdown = 0 * time.Millisecond
                w.Write([]byte(`{ "text": "Made server fast again; virtually no delay"}`))
        })

        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatal("ListenAndServe:", err)
        }
}
