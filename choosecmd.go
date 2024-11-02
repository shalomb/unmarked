package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(chooseMarksCmd)
}

// markCmd represents the mark command
var chooseMarksCmd = &cobra.Command{
	Use:   "choose",
	Short: "Choose from available marked windows",
	Long:  `Choose from all the marked windows`,
	Run: func(_ *cobra.Command, _ []string) {
		log.Printf("choosing marks under %v", stateHome)
		chooseMarks()
	},
}

func chooseMarks() {
	matches, _ := filepath.Glob(fmt.Sprintf("%s/*", stateHome))
	lines := []string{}
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

			line := fmt.Sprintf("%s -> %s [%s]\n", filepath.Base(match), title, app)
			lines = append(lines, line)
		}
	}

	chooseCmd := exec.Command(
		"choose", "-n", "20", "-s", "20", "-b", "fac898",
		"-c", "FF7518", "-p", "Choose a mark")
	chooseIn, _ := chooseCmd.StdinPipe()
	chooseOut, _ := chooseCmd.StdoutPipe()

	chooseCmd.Start()
	chooseIn.Write([]byte(strings.Join(lines, "\n")))
	chooseIn.Close()

	chooseBytes, _ := io.ReadAll(chooseOut)
	chooseCmd.Wait()

	choice := string(chooseBytes)
	if len(choice) == 0 {
		log.Fatalf("No/empty response from choose, aborting", choice)
	}

	mark := string(choice[0])
	if len(mark) == 0 {
		log.Fatalf("Error deriving mark (%v) from choose (%v)", mark, choice)
	}

	m := NewWinMarker()
	m.SummonMark(mark)
}
