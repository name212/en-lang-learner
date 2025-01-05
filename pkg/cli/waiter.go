package cli

import (
	"bufio"
	"os"
	"strings"
)

type AnswerWaiter struct{}

func NewAnswerWaiter() *AnswerWaiter {
	return &AnswerWaiter{}
}

func (t *AnswerWaiter) WaitAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), nil
}
