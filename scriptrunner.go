package main

import (
	"bytes"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// ShellScript is ..
var scriptRunner *Script

// Script is ...
type Script struct {
	workDir string
	script  string
	args    []string
}

// Script is ..
func (s *Script) Script() string {
	return s.script
}

// Args is ..
func (s *Script) Args() []string {
	return s.args
}

// Exec is ..
func (s *Script) Exec() (int, bytes.Buffer, bytes.Buffer, error) {
	e := s.script
	a := s.args
	log.Printf("Running [%v] with [%v]", e, a)

	var err error

	tmpfile, err := os.CreateTemp(os.TempDir(), "unmarked.*.script")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(e)); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(tmpfile.Name(), 0700)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(tmpfile.Name(), a...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr

	err = cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Errorf("exitcode: %d: %s (%s)", exiterr.ExitCode(), err, stderr.String())
			log.Printf("\a")
			return exiterr.ExitCode(), stdout, stderr, err
		}
	}

	return 0, stdout, stderr, err
}
