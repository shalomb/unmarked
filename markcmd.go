package main

import (
	"log"

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
		var mark = args[0]
		log.Printf("Marking %v", mark)
		m := NewWinMarker()
		m.MarkWindow([]rune(mark)[0])
		// m.RaisePreviousWindow()
	},
}
