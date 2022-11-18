package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type ScriptRunner struct {
	script string
}

func NewScriptRunner() *ScriptRunner {
	log.Printf("Setting up script runner")
	return new(ScriptRunner)
}

func (r *ScriptRunner) run(s *OSAScript, a ...string) (int, bytes.Buffer, bytes.Buffer, error) {
	log.Printf("Running [%v] with [%v]", s, a)

	var err error

	tmpfile, err := ioutil.TempFile(os.TempDir(), "unmarked.*.script")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(s.script)); err != nil {
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
