package main

import (
    "embed"
    "io/fs"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "tempmail-app/application"
)

//go:embed gui/*
var gui embed.FS

func main() {
    args := os.Args[1:]

    if len(args) > 0 && args[0] == "dev" {
        u, err := url.Parse("http://localhost:5500/")
        if err != nil {
            log.Fatal(err)
        }
        http.Handle("/", httputil.NewSingleHostReverseProxy(u))
    } else {
        pub, err := fs.Sub(gui, "gui")
        if err != nil {
            log.Fatal(err)
        }
        http.Handle("/", http.FileServer(http.FS(pub)))
    }

    application.Start()
}
