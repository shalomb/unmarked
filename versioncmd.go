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
	Short: "Print version information",
	Long:  `Print the version and build information`,
	Run: func(_ *cobra.Command, _ []string) {
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
