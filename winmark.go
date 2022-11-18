package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
)

type WinMarker struct {
	stateHome string
}

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

func (w *WinMarker) SummonMark(args []string) {
	log.Printf("SummonMark called: %v", args)
	mark := args[0]
	markFile := path.Join(w.stateHome, mark)
	if _, err := os.Stat(markFile); errors.Is(err, os.ErrNotExist) {
		log.Warnf("Mark file does not exist: %v (%v)", mark, markFile)
	} else {
		log.Printf("Summoning: %v", mark)
		if data, err := os.ReadFile(markFile); err == nil {
			s := strings.Split(string(data), ",")
			log.Printf("got: %v", s)

			if _, stdout, _, err := osascript(heredoc.Doc(`
				on run argv
					set processName to ""
					set windowTitle to ""

					set processName to item 1 of argv
					try
						set windowTitle to item 2 of argv
					end try

					tell application "System Events" to tell process processName
						set frontmost to true
						windows where title contains windowTitle
						if result is not {} then
							perform action "AXRaise" of item 1 of result
						end if
						# activate
						reopen
						set visible to true

					end tell
				end run
			`), s...); err == nil {
				log.Debugf("We got good: %v", stdout.String())
			}
		}
	}
}

func (w *WinMarker) RaisePreviousWindow() {
	log.Printf("Marking : %v", w)
	if _, stdout, _, err := osascript(heredoc.Doc(`
		tell application "System Events"
			keystroke tab using command down
		end tell
	`)); err == nil {
		log.Debugf("We got good: %v", stdout.String())
	}
}

func (w *WinMarker) MarkWindow(r rune) {
	log.Printf("Marking : %v", string(r))
	if _, stdout, _, err := osascript(heredoc.Doc(`
		tell application "System Events"
			set activeApp to (first application process whose frontmost is true)
			set activeAppName to name of activeApp
			set frontWin to first window of activeApp
		end tell
		return {activeAppName, frontWin}
	`)); err == nil {
		log.Debugf("We got good: %v", stdout.String())
		w.SaveMark(r, stdout.Bytes())
	}
}

func (w *WinMarker) SaveMark(r rune, b []byte) {
	log.Printf("Saving mark: %v", r)
	target := path.Join(w.stateHome, string(r))
	if err := os.WriteFile(target, b, 0600); err == nil {
		log.Printf("Mark written: %v -> (%v)", r, string(b))
	} else {
		log.Errorf("Error writing mark: %v -> (%v)", r, string(b))
	}
}
