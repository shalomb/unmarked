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
	if err := Execute(); err != nil {
		log.Fatalf("Error running unmarked: %v", err)
		os.Exit(3)
	}
}

var (
	appName   string = "unmarked"
	stateHome string = path.Join(xdg.StateHome, appName)
	Version   string
	BuildTime string
	BuildHost string
	Branch    string
)
