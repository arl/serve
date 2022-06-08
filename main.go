package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	defaultRoot = "."

	defaultHost     = "localhost"
	defaultPort     = "8000"
	defaultHostPort = defaultHost + ":" + defaultPort
)

const helpTxt = `Serves the content of a directory as HTTP
usage: serve [-h] [dir] [addr]
parameters:
    dir         File server root directory, defaults to current directory
    addr        Address to listen on. Defaults to "` + defaultHostPort + `"
                Format is host:port where "host" can be omitted to listen on all
                network interfaces.
    -h/--help   Show this message`

func main() {
	log.SetFlags(0)
	log.SetPrefix("[serve] ")
	root, addr, help, exit := parseCommandLine(os.Args[1:])

	if help {
		fmt.Fprintln(os.Stderr, helpTxt)
		os.Exit(exit)
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(2)
	}

	if host == "" {
		host = "*"
	}

	log.Println("serving", root, "on http://"+host+":"+port+"/")
	panic(http.ListenAndServe(addr, noCache(http.FileServer(http.Dir(root)))))
}

// parseCommandLine
func parseCommandLine(args []string) (root, addr string, help bool, exit int) {
	root = defaultRoot
	addr = defaultHostPort

	nargs := len(args)

	switch {
	case nargs == 0: // all default
	case args[0] == "-h", args[0] == "--help":
		help = true
	case nargs > 2:
		help = true
		exit = 1
	case nargs == 1: // just dir
		root = args[0]
	default:
		root = args[0]
		addr = args[1]
	}
	return
}
