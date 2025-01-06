package cli

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/name212/en-lang-learner/pkg"
)

type AnswerWaiter struct{}

func NewAnswerWaiter() *AnswerWaiter {
	return &AnswerWaiter{}
}

func (t *AnswerWaiter) WaitAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", pkg.ErrStopByUser
		}

		return "", err
	}

	return strings.TrimSpace(line), nil
}
