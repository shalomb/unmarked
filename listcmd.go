package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// "strings"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listMarksCmd)
}

// markCmd represents the mark command
var listMarksCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List marked windows",
	Long:    `List all marked windows, displaying their mark with the window title and app`,
	Run: func(_ *cobra.Command, _ []string) {
		listMarks()
	},
}

func listMarks() {
	log.Printf("Listing marks under %v", stateHome)
	currentMarks, _ := findMarks()
	for _, v := range currentMarks {
		fmt.Printf("%v", v)
	}
}

func findMarks() ([]string, []string) {
	windowMap := getWindows()
	matches, _ := filepath.Glob(fmt.Sprintf("%s/*", stateHome))
	lines := []string{}
	stale := []string{}
	for _, match := range matches {
		f, _ := os.Stat(match)
		if !f.IsDir() {
			data, err := os.ReadFile(match)
			if err != nil {
				log.Printf("ERROR: could not read file", match)
			}

			id, err := jq(".id", string(data))
			if err != nil {
				log.Printf("Could not find .id under %v", data)
			}

			if obj, ok := windowMap[fmt.Sprint(id)]; ok {
				line := fmt.Sprintf("%v ⟶ %v\n",
					filepath.Base(match),
					strings.Join(obj, " | "))
				lines = append(lines, line)
			} else {
				log.Printf("id %v not in windowMap", id)
				stale = append(stale, match)
			}
		}
	}
	return lines, stale
}

func getWindows() map[string][]string {
	windowMap := make(map[string][]string)

	sc := yabaiscript(heredoc.Doc(`yabai -m query --windows`))
	if _, stdout, stderr, err := sc.Exec(); err == nil {

		var data []interface{}
		if err := json.Unmarshal([]byte(stdout.String()), &data); err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		for _, v := range data {
			winMap := v.(map[string]interface{})
			id := fmt.Sprint(winMap["id"])
			title := winMap["title"].(string)
			app := winMap["app"].(string)
			windowMap[id] = []string{title, id, app}
		}
	} else {
		log.Fatalf("Error running yabai: err: %s, stderr: %s", err, stderr.String())
	}

	return windowMap
}
