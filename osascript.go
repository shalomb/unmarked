package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
)

func osascript(s string, args ...string) *Script {
	script := fmt.Sprintf(
		heredoc.Doc(`
		#!/usr/bin/osascript
		%s
	`), s)
	return &Script{
		script: script,
		args:   args,
	}
}
