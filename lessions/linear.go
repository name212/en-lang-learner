package lessions

import "github.com/name212/en-lang-learner/pkg"

type linearLession struct {
	tasks         []pkg.Task
	nextTaskIndex int
}

func newLinearLession(tasks []pkg.Task) *linearLession {
	return &linearLession{
		tasks:         tasks,
		nextTaskIndex: 0,
	}
}

func (l *linearLession) NextTask() (pkg.Task, error) {
	if l.nextTaskIndex >= len(l.tasks) {
		return nil, pkg.ErrNoNextTask
	}

	t := l.tasks[l.nextTaskIndex]
	l.nextTaskIndex++

	return t, nil
}
