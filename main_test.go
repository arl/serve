package main

import (
	"strings"
	"testing"
)

func TestParseCommandLine(t *testing.T) {
	tests := []struct {
		args     []string
		wantRoot string
		wantAddr string
		wantHelp bool
		wantExit int
	}{
		{args: []string{"-h"}, wantHelp: true, wantExit: 0},
		{args: []string{"a", "b", "c"}, wantHelp: true, wantExit: 1},
		{args: nil, wantAddr: defaultHostPort, wantRoot: "."},
		{args: []string{"/foo/bar"}, wantRoot: "/foo/bar", wantAddr: defaultHostPort},
		{args: []string{"/foo/bar", "host:port"}, wantRoot: "/foo/bar", wantAddr: "host:port"},
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
				t.Errorf("parseCommandLine(`%v`) gotHelp = %v, want %v", cli, gotHelp, tt.wantHelp)
			}

			if gotHelp {
				return
			}

			if gotAddr != tt.wantAddr {
				t.Errorf("parseCommandLine(`%v`) gotAddr = %v, want %v", cli, gotAddr, tt.wantAddr)
			}

			if gotRoot != tt.wantRoot {
				t.Errorf("parseCommandLine(`%v`) gotRoot = %v, want %v", cli, gotRoot, tt.wantRoot)
			}
		})
	}
}
