package main

import (
	"strings"
	"testing"
)

func TestParseCommandLine(t *testing.T) {
	tests := []struct {
		args []string

		wantAddr string
		wantDir  string
		wantHelp bool
		wantExit int
	}{
		{args: []string{"-h"}, wantHelp: true, wantExit: 0},
		{args: []string{"a", "b", "c"}, wantHelp: true, wantExit: 1},
		{args: nil, wantAddr: defaultHostPort, wantDir: "."},
		{args: []string{":80"}, wantAddr: ":80", wantDir: "."},
		{args: []string{"0.0.0.0:4567", "/foo/bar"}, wantAddr: "0.0.0.0:4567", wantDir: "/foo/bar"},
	}

	for _, tt := range tests {
		args := []string{"serve"}
		for _, a := range tt.args {
			args = append(args, `"`+a+`"`)
		}

		cli := strings.Join(args, " ")
		t.Run(cli, func(t *testing.T) {
			gotRoot, gotAddr, gotHelp, gotExit := parseCommandLine(tt.args)
			_ = gotExit

			if gotHelp != tt.wantHelp {
				t.Fatalf("parseCommandLine(`%v`) gotHelp = %v, want %v", cli, gotHelp, tt.wantHelp)
			}

			if gotHelp {
				return
			}

			if gotAddr != tt.wantAddr {
				t.Fatalf("parseCommandLine(`%v`) gotAddr = %v, want %v", cli, gotAddr, tt.wantAddr)
			}

			if gotRoot != tt.wantDir {
				t.Fatalf("parseCommandLine(`%v`) gotRoot = %v, want %v", cli, gotRoot, tt.wantDir)
			}
		})
	}
}
