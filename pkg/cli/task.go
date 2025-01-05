package cli

import (
	"bufio"
	"os"

	"github.com/name212/en-lang-learner/pkg"
)

type CLIInputTaks struct {
	*pkg.Task
}

func NewCLIInputTask(target string, answer string) *CLIInputTaks {
	return &CLIInputTaks{
		&pkg.Task{
			AnswerStr: answer,
			TargetStr: target,
		},
	}
}

func (t *CLIInputTaks) WaitAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return line, nil
}
