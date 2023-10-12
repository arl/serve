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
    addr        Address to listen on. Defaults to "` + defaultHostPort + `"
                Format is host:port where "host" can be omitted to listen on all
                network interfaces.
    dir         File server root directory, defaults to current directory
    -h/--help   Show this message`

func main() {
	log.SetFlags(0)
	log.SetPrefix("[serve] ")

	dir, addr, help, exit := parseCommandLine(os.Args[1:])
	if help {
		fmt.Fprintln(os.Stderr, helpTxt)
		os.Exit(exit)
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if host == "" {
		host = "*"
	}

	if _, err := os.Stat(dir); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	log.Println("serving", dir, "on http://"+host+":"+port+"/")
	panic(http.ListenAndServe(addr, noCache(http.FileServer(http.Dir(dir)))))
}

// parseCommandLine
func parseCommandLine(args []string) (dir, addr string, help bool, exit int) {
	dir = defaultRoot
	addr = defaultHostPort

	nargs := len(args)

	switch {
	case nargs == 0: // all default
	case args[0] == "-h", args[0] == "--help":
		help = true
	case nargs > 2:
		help = true
		exit = 1
	case nargs == 1: // just addr
		addr = args[0]
	default:
		addr = args[0]
		dir = args[1]
	}
	return
}
