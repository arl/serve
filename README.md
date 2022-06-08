# Serve

## Installation


```
go install github.com/arl/serve@latest
```

## Usage

Like `python -m SimpleHTTPServer` but _simpler_.

    $ serve -h
    Serves the content of a directory as HTTP
    usage: serve [-h] [root] [[host]:port]
        command summary
            [root]         File server root directory, defaults to current directory
            [[host]:port]  Address to listen on, defaults to "localhost:8000"
                           "host" can be omitted to listen on all network interfaces.
            -h/--help      This help text


Examples:

* serve the current directory on localhost and default port (8000)

```
$ serve
```


* serve `/foo/bar` on localhost and default port (8000)

```
$ serve /foo/bar
```


- serve the current directory on localhost and port 7777

```
$ serve /foo/bar :7777
```


- serve the current directory on localhost and port 7777

```
$ serve . :7777
```
