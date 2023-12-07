# Serve

Like `python -m SimpleHTTPServer` but even _simpler_!

## Installation

    go install github.com/arl/serve@latest


## Usage

    $ serve -h
    Serves the content of a directory as HTTP
    usage: serve [-h] [addr] [dir]
    options:
        addr        Address to listen on. Defaults to "localhost:8000"
                    Format is host:port where "host" can be omitted to listen on all
                    network interfaces.
        dir         File server root directory, defaults to current directory
        -h/--help   Show this message

Once started, `serve` logs HTTP requests as follows:

    $ serve
    [serve] serving . on http://localhost:8000/
    [serve] "GET / HTTP/1.1" 0 7 58.398µs
    [serve] "GET /index.html HTTP/1.1" 200 1592 3.958531ms
    [serve] "GET /main.css HTTP/1.1" 200 821 109.286µs

For example, the last line indicates that an `HTTP/1.1` request to `main.css` has been served, it ended with a 200 status, 821 bytes were served and that took 109µs.


## Examples

* Serves current directory on defaut address `http://localhost:8000`

```
serve
```

* Serves current directory on `http://*:80` (all network interfaces)

```
serve :80
```

* Serves `/foo/bar` directory on `http://*:7777` (all network interfaces)

```
serve :7777 /foo/bar
```
