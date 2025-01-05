package lession

import (
	"math/rand/v2"

	"github.com/name212/en-lang-learner/pkg"
)

// BiderictionalRandomLession
// Get task list by default all targets in english and answers on russian
// Next create new lists with en -> ru and ru -> en tasks
// joined list and sort randomly
// will be finished will return ErrNoNextTask for stopping lession
type BiderictionalRandomLession struct {
	*LinearLession
}

func extendAndRandomTasks(tasks []pkg.Task) []pkg.Task {
	newTasks := make([]pkg.Task, 0, 2*len(tasks))
	for _, t := range tasks {
		newTasks = append(newTasks, pkg.NewBaseTask(t.Target(), t.Answer()))
		newTasks = append(newTasks, pkg.NewBaseTask(t.Answer(), t.Target()))
	}

	rand.Shuffle(len(newTasks), func(i, j int) {
		newTasks[i], newTasks[j] = newTasks[j], newTasks[i]
	})

	return newTasks
}

func NewBiderictionalRandomLession(tasks []pkg.Task) *BiderictionalRandomLession {
	randomizeTasks := extendAndRandomTasks(tasks)
	return &BiderictionalRandomLession{
		LinearLession: &LinearLession{
			tasks: randomizeTasks,
		},
	}
}
