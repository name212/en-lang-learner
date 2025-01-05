package pkg

import "errors"

var ErrNoNextTask = errors.New("Do not have tasks")

type Task interface {
	Target() string
	Answer() string
}

type Lession interface {
	NextTask() (Task, error)
}

type BaseTask struct {
	TargetStr string
	AnswerStr string
}

func NewBaseTask(target string, answer string) *BaseTask {
	return &BaseTask{
		TargetStr: target,
		AnswerStr: answer,
	}
}

func (t *BaseTask) Target() string {
	return t.TargetStr
}

func (t *BaseTask) Answer() string {
	return t.AnswerStr
}
