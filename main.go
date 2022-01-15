package main

import (
	"flag"
	"strconv"
	"net/http"
)

func main() {
    port := flag.Int("port", 8080, "port number")
    flag.Parse()

    http.ListenAndServe(":" + strconv.Itoa(*port), http.FileServer(http.Dir(flag.Arg(0))))
}
