package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(listMarksCmd)
}

// markCmd represents the mark command
var listMarksCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the marked windows",
	Long:  `List all marked windows, displaying their mark as well as the extended titles`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Listing marks under %v", stateHome)
		listMarks()
	},
}

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
