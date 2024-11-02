package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
)

func yabaiscript(s string, args ...string) *Script {
	return &Script{
		args: args,
		script: fmt.Sprintf(heredoc.Doc(`
		#!/usr/bin/env bash

		set -o errexit -o nounset -o noclobber -o pipefail
		shopt -s extglob nullglob

		[[ ${DEBUG-} ]] && set -xv

		%s
	`), s),
	}
}
