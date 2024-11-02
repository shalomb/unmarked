package main

import (
	"fmt"
	"os"

	// "strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanMarksCmd)
}

// markCmd represents the mark command
var cleanMarksCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"purge"},
	Short:   "Remove stale marks",
	Long:    `Remove marks that are stale and do not point at any windows`,
	Run: func(_ *cobra.Command, _ []string) {
		cleanMarks()
	},
}

func cleanMarks() {
	log.Printf("Listing marks under %v", stateHome)
	_, staleMarks := findMarks()
	for _, v := range staleMarks {
		fmt.Printf("rm '%v'\n", v)
		if err := os.Remove(v); err != nil {
			log.Warningf("Failure removing mark file: %v, %v", v, err)
		}
	}
}
