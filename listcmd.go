package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
