package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summonCmd)
}

// summonCmd represents the summon command
var summonCmd = &cobra.Command{
	Use:   "summon",
	Short: "Summon a given window by its mark",
	Long:  `Summon a given window by its mark`,
	Run: func(cmd *cobra.Command, args []string) {
		var mark string
		if len(args) > 0 {
			mark = args[0]
		} else {
			mark = "@"
		}
		log.Printf("summon called: %v", mark)
		m := NewWinMarker()
		m.SummonMark(mark)
	},
}
