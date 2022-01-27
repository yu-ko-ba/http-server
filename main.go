package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        log.Fatalln(err)
    }

    for _, addr := range addrs {
        a, ok := addr.(*net.IPNet)
        if ok && a.IP.To4() != nil && !strings.Contains(a.String(), "127.0.0.1") {
            fmt.Println("IP address : " + a.String())
        }
    }

    portNumber := flag.Int("port", 8080, "port number")
    flag.Parse()

    port := strconv.Itoa(*portNumber)

    fmt.Println("Port number: " + port)

    http.ListenAndServe(":" + port, http.FileServer(http.Dir(flag.Arg(0))))
}
