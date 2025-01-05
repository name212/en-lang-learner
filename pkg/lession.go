package pkg

import "errors"

var ErrNoNextTask = errors.New("Do not have tasks")

type TaskRunner interface {
	WaitAnswer() (string, error)
	Target() string
	Answer() string
}

type Lession interface {
	NextTask() (TaskRunner, error)
}

type Task struct {
	TargetStr string
	AnswerStr string
}

func (t *Task) Target() string {
	return t.TargetStr
}

func (t *Task) Answer() string {
	return t.AnswerStr
}
