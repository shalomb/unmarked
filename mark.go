package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
)

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

// SummonMark is ...
func (w *WinMarker) SummonMark(mark string) {
	markFile := path.Join(w.stateHome, mark)
	log.Printf("SummonMark called: %+v, %+v", mark, markFile)
	if _, err := os.Stat(markFile); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Mark file does not exist: %v (%v)", mark, markFile)
	}

	if data, err := os.ReadFile(markFile); err == nil {
		log.Printf("got: %v", string(data))

		winid, err := jq(".id", string(data))
		if err != nil {
			log.Printf("Error reading .id from markfile (%v)", markFile)
		}

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

		log.Printf("Marking window with id %v with %v", id, string(r))
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
		log.Printf("Mark written: %v -> (%v)", string(r), string(b))
	} else {
		log.Errorf("Error writing mark: %v -> (%v)", r, string(b))
	}
	return nil
}
