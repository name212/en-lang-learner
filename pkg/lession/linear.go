package lession

import "github.com/name212/en-lang-learner/pkg"

// LinearLession
// Get task list and take task in passed order
// after end of task list will return ErrNoNextTask for stopping lession
type LinearLession struct {
	tasks         []pkg.Task
	nextTaskIndex int
}

func NewLinearLession(tasks []pkg.Task) *LinearLession {
	return &LinearLession{
		tasks:         tasks,
		nextTaskIndex: 0,
	}
}

func (l *LinearLession) NextTask() (pkg.Task, error) {
	if l.nextTaskIndex >= len(l.tasks) {
		return nil, pkg.ErrNoNextTask
	}

	t := l.tasks[l.nextTaskIndex]
	l.nextTaskIndex++

	return t, nil
}
