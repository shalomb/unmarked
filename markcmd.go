package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markCmd)
}

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Mark a given window with a letter or number",
	Long: `Windows can be marked and assigned letters or numbers as
	shortcuts that can later be used in activating/showing those windows`,
	Run: func(cmd *cobra.Command, args []string) {
		var mark string
		if len(args) > 0 {
			mark = args[0]
		} else {
			mark = "@"
		}
		log.Printf("Marking current window with %v", mark)
		m := NewWinMarker()
		m.MarkWindow([]rune(mark)[0])
	},
}
