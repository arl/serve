package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

const (
	defaultRoot = "."

	defaultHost = "localhost"
	defaultPort = "8000"
	defaultAddr = defaultHost + ":" + defaultPort
)

const helpTxt = `Serves the content of a directory as HTTP
usage: serve [-h] [root] [[host]:port]
    command summary
        [root]         File server root directory, defaults to current directory
        [[host]:port]  Address to listen on, defaults to "` + defaultAddr + `"
                       "host" can be omitted to listen on all network interfaces.
        -h/--help      This help text
`

func main() {
	root, addr, help := parseCommandLine(os.Args[1:])

	if help {
		fmt.Fprintln(os.Stderr, helpTxt)
		os.Exit(1)
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	if host == "" {
		host = "*"
	}

	fmt.Println("serving", root, "on http://"+host+":"+port+"/")
	panic(http.ListenAndServe(addr, noCache(http.FileServer(http.Dir(root)))))
}

// parseCommandLine
func parseCommandLine(args []string) (root, addr string, help bool) {
	nargs := len(args)
	if nargs == 0 {
		return defaultRoot, defaultAddr, false
	}

	if args[0] == "-h" || args[0] == "--help" || nargs > 2 {
		help = true
		return "", "", help
	}

	if nargs == 1 {
		return args[0], defaultAddr, false
	}

	return args[0], args[1], false
}
