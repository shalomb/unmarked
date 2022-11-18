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
		log.Printf("summon called: %v", args)
		m := NewWinMarker()
		m.SummonMark(args)
	},
}
