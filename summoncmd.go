package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summonCmd)
}

// summonCmd represents the summon command
var summonCmd = &cobra.Command{
	Use:   "summon",
	Short: "Summon a window by mark",
	Long:  `Summon a given window by its mark`,
	Run: func(_ *cobra.Command, args []string) {
		var mark string

		if len(args) > 0 {
			mark = args[0]
		} else {
			mark = "@"
		}

		m := NewWinMarker()
		m.SummonMark(mark)
	},
}

// SummonMark is ...
func (w *WinMarker) SummonMark(mark string) {
	markFile := path.Join(w.stateHome, mark)
	log.Printf("SummonMark called: %+v, %+v", mark, markFile)

	if _, err := os.Stat(markFile); errors.Is(err, os.ErrNotExist) {
		osascript(fmt.Sprintf((heredoc.Doc(`
            display alert "No mark defined for '%s'" giving up after 1.5
		`)), mark))
		log.Fatalf("Mark file does not exist: %v (%v)", mark, markFile)
	}

	if data, err := os.ReadFile(markFile); err == nil {
		log.Debugf("Mark data: %v", string(data))

		winid, err := jq(".id", string(data))
		if err != nil {
			log.Printf("Error reading .id from markfile (%v)", markFile)
		}

		title, err := jq(".title", string(data))
		app, err := jq(".app", string(data))
		log.Printf("Window Info: %v, %v, %v", winid, title, app)

		sc := yabaiscript(heredoc.Doc(`
				yabai -m window --focus  $1
			`), fmt.Sprintf("%.0f", winid))
		if ec, stdout, stderr, err := sc.Exec(); err == nil {
			log.Printf("OK: %v, %+v", ec, stdout.String())
		} else {
			log.Fatalf("Error running yabai: err: %s, stderr: %s", err, stderr.String())
		}

	}
}

// RaisePreviousWindow is ...
func (w *WinMarker) RaisePreviousWindow() {
	log.Printf("Raising previous: %v", w)
	sc := yabaiscript(heredoc.Doc(`
				yabai -m window --focus prev
			`))
	if ec, stdout, stderr, err := sc.Exec(); err == nil {
		log.Printf("OK: %v, %+v", ec, stdout.String())
	} else {
		log.Fatalf("Error: err: %s, stderr: %s", err, stderr.String())
	}
}
