package tts

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"reflect"
)

type runner struct {
	stdin io.Reader
}

func newRunner() *runner {
	return &runner{}
}

func (r *runner) exec(cmdName string, args ...string) ([]byte, error) {
	cmd := exec.Command(cmdName, args...)
	stdoutBuf := &bytes.Buffer{}
	stderrBuf := &bytes.Buffer{}
	cmd.Stdout = stdoutBuf
	cmd.Stderr = stderrBuf

	if r.stdin != nil {
		refVal := reflect.ValueOf(r.stdin)
		if !refVal.IsNil() {
			cmd.Stdin = r.stdin
		}
	}

	err := cmd.Run()
	if err != nil {
		code := -1
		if errExit, ok := err.(*exec.ExitError); ok {
			code = errExit.ExitCode()
		}

		err = fmt.Errorf("exitcode %d; %s\n%s", code, err.Error(), stderrBuf.String())
		return nil, err
	}

	return stdoutBuf.Bytes(), nil
}
