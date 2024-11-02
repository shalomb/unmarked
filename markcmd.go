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
	rootCmd.AddCommand(markCmd)
}

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Mark the active window with a letter/number",
	Long: `Windows can be marked and assigned letters or numbers as
	shortcuts that can later be used in activating/showing those windows`,
	Run: func(_ *cobra.Command, args []string) {
		var mark string

		if len(args) > 0 {
			mark = args[0]
		} else {
			mark = "@"
		}

		m := NewWinMarker()
		m.MarkWindow([]rune(mark)[0])
	},
}

// WinMark is ...
type WinMark struct {
	id   string
	data interface{}
}

// WinMarker is ...
type WinMarker struct {
	stateHome string
}

// NewWinMarker is ...
func NewWinMarker() *WinMarker {
	m := new(WinMarker)
	m.stateHome = stateHome
	if _, err := os.Stat(m.stateHome); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(m.stateHome, 0700); err != nil {
			log.Errorf("Error creating stateHome (%v): %v", m.stateHome, err)
		}
	}
	return m
}

// MarkWindow is ...
func (w *WinMarker) MarkWindow(r rune) {
	s := yabaiscript(heredoc.Doc(`
		yabai -m query --windows --window
	`))
	if _, stdout, _, err := s.Exec(); err == nil {
		log.Debugf("We got good: %v", stdout.String())

		id, err := jq(".id", stdout.String())
		if err != nil {
			log.Fatalf("Error getting .id, %v", err)
		}

		log.Printf("Marking window with id %v mark %v", id, string(r))
		if err := w.SaveMark(r, stdout.Bytes()); err != nil {
			log.Fatalf("Error saving mark, %v", err)
		}

		title, err := jq(".title", stdout.String())
		if err != nil {
			log.Fatalf("Error getting .id, %v", err)
		}

		n := osascript(
			fmt.Sprintf((heredoc.Doc(`
					display dialog "%s -> %s" giving up after 2
		`)), string(r), title.(string)))
		if _, _, _, err := n.Exec(); err != nil {
			log.Fatalf("Failed invoking notification, %v", err)
		}
	}
}

// SaveMark is ...
func (w *WinMarker) SaveMark(r rune, b []byte) error {
	target := path.Join(w.stateHome, string(r))

	log.Printf("Saving mark '%+v' to '%v'", string(r), target)
	if err := os.WriteFile(target, b, 0600); err == nil {
		log.Debugf("Mark written: %v -> (%v)", string(r), string(b))
	} else {
		log.Errorf("Error writing mark: %v -> (%v)", r, string(b))
	}

	return nil
}
