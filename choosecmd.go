package main

import (
	"io"
	"os/exec"
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
	Short: "Choose from marked windows",
	Long:  `Choose from all the marked windows`,
	Run: func(_ *cobra.Command, _ []string) {
		if !commandExists("choose") {
			log.Fatalf("tool 'choose' is not installed. Install with 'brew install choose-gui'")
		}

		log.Printf("choosing marks under %v", stateHome)
		chooseMarks()
	},
}

func chooseMarks() {
	lines, _ := findMarks()

	chooseCmd := exec.Command(
		"choose", "-n", "20", "-s", "20", "-b", "fac898",
		"-c", "FF7518", "-p", "Choose mark")
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
