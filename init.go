package main

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func init() {
	for _, tool := range []string{"ls", "yabai", "skhd"} {
		log.Debugf("tool exists: %v -> %v", tool, commandExists(tool))
		if !commandExists(tool) {
			log.Fatalf("tool %v is not installed. Install with 'brew install %v'", tool, tool)
		}
	}
}

// as util
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
