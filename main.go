/*
Copyright © 2022 Shalom Bhooshi
*/

package main

import (
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
)

func main() {
	if err := InitCobra(); err != nil {
		log.Fatalf("Error running unmarked: %v", err)
		os.Exit(3)
	}
}

var (
	AppName   string
	Branch    string
	BuildHost string
	BuildTime string
	GoVersion string
	GoArch    string
	GoOS      string
	Version   string

	stateHome string = path.Join(xdg.StateHome, AppName)
)
