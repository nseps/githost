package main

import (
	"net/http"
	"golang.org/x/net/webdav"
	"flag"
	"strconv"
	"fmt"
	"os"
	"./lib"
)

var (
	host     = flag.String("host", "localhost", "Hostname")
	port     = flag.Int("port", 80, "Port")
	repoPath = flag.String("repoPath", ".", "Repository root folder")
	cert     = flag.String("cert", "cert.pem", "TLS Certificate")
	key      = flag.String("key", "key.pem", "TLS Key")
)

func main() {

	var c githost.Config = "config"

	c.PutWithURL("http", "www.ena.com", "enabled", "true")
	c.PutWithURL("http", "www.dio.com", "enabled", "false")
	c.PutWithURL("http", "www.tria.com", "enabled", "true")

	val, err := c.GetInt("tria", "int")
	if err != nil {
		fmt.Println("Erroou", err)
	}

	fmt.Println("VAL:", val)

	os.Exit(0)

	handler := webdav.Handler{
		FileSystem: webdav.Dir(*repoPath),
		LockSystem: webdav.NewMemLS(),
	}
	http.Handle("/", &handler)
	http.ListenAndServe( *host + ":" + strconv.Itoa(*port), nil)
}
