package main

import "bytes"

type OSAScript struct {
	script string
}

var (
	scriptRunner = NewScriptRunner()
)

func osascript(s string, a ...string) (int, bytes.Buffer, bytes.Buffer, error) {
	s = "#!/usr/bin/env osascript\n" + s
	return scriptRunner.run(&OSAScript{script: s}, a...)
}
