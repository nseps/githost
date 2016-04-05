package main

import (
	"net/http"
	"golang.org/x/net/webdav"
)

func main() {

	handler := webdav.Handler{
		FileSystem: webdav.Dir("/home/nikolas/repos"),
		LockSystem: webdav.NewMemLS(),
	}
	http.Handle("/", &handler)
	http.ListenAndServe("localhost:80", nil)
}
