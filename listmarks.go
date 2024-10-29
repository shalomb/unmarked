package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func listMarks() {
	matches, _ := filepath.Glob(fmt.Sprintf("%s/*", stateHome))
	for _, match := range matches {
		f, _ := os.Stat(match)
		if !f.IsDir() {
			data, err := os.ReadFile(match)
			if err != nil {
				log.Printf("ERROR: could not read file", match)
			}

			app, err := jq(".app", string(data))
			if err != nil {
				log.Printf("Could not find .app under %v", data)
			}

			title, err := jq(".title", string(data))
			if err != nil {
				log.Printf("Could not find .title under %v", data)
			}

			fmt.Printf("%s -> %s [%s]\n", filepath.Base(match), title, app)
		}
	}
}
