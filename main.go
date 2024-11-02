/*
Copyright © 2022 Shalom Bhooshi
*/

// main is ...
package main

import (
	"log"
	"path"

	"github.com/adrg/xdg"
)

func main() {
	if err := InitCobra(); err != nil {
		log.Fatalf("Error running unmarked: %v", err)
	}
}

var (
	// AppName is the application name set via GOLDFLAGS
	AppName string
	// Branch is the git branch name set via GOLDFLAGS
	Branch string
	// BuildHost is the build hostname set via GOLDFLAGS
	BuildHost string
	// BuildTime is the application build time set via GOLDFLAGS
	BuildTime string
	// GoVersion is the Golang version set via GOLDFLAGS
	GoVersion string
	// GoArch is the go architecture set via GOLDFLAGS
	GoArch string
	// GoOS is the go OS name set via GOLDFLAGS
	GoOS string
	// Version is the application version string set via GOLDFLAGS
	Version string

	stateHome string = path.Join(xdg.StateHome, AppName)
)
