package main

import (
	"log"
	"net/http"
	"time"
)

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

func noCache(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Delete any ETag headers that may have been set
		for _, v := range etagHeaders {
			if r.Header.Get(v) != "" {
				r.Header.Del(v)
			}
		}

		rw := &responseWriter{ResponseWriter: w}

		// Set our NoCache headers
		for k, v := range noCacheHeaders {
			rw.Header().Set(k, v)
		}

		start := time.Now()
		h.ServeHTTP(rw, r)
		dur := time.Since(start)

		log.Printf(`"%s %s %s" %d %d %s`, r.Method, r.RequestURI, r.Proto, rw.status, rw.written, dur)
	}
}

type responseWriter struct {
	http.ResponseWriter

	status  int
	written int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.written = n
	return n, err
}
