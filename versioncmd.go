package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(heredoc.Doc(`
      %s version %s
      %s, %s from %s on %s
      go version %s %s/%s
		`), AppName, Version,
			BuildTime, Version,
			Branch, BuildHost,
			GoVersion, GoOS, GoArch)
	},
}
