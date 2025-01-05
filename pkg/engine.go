package pkg

import (
	"errors"
	"strings"
)

type TextToSpeech interface {
	Say(text string) error
}

type Outputer interface {
	OutTarget(t string) error
	OutAnswer(a string, isTrue bool) error
	OutPrompt() error
}

type AnswerWaiter interface {
	WaitAnswer() (string, error)
}

type Stats struct {
	TotalTasks int
	OkTaksks   int
}

type Engine struct {
	lessionLenght int
	stats         *Stats

	tts      TextToSpeech
	lession  Lession
	outputer Outputer
	waiter   AnswerWaiter
}

func NewEngine(lession Lession, tts TextToSpeech, outputer Outputer, waiter AnswerWaiter) *Engine {
	return &Engine{
		tts:      tts,
		lession:  lession,
		outputer: outputer,
		waiter:   waiter,

		stats: &Stats{},
	}
}

func (e *Engine) WithLessionLenght(l int) *Engine {
	e.lessionLenght = l
	return e
}

func (e *Engine) Run() (*Stats, error) {
	for {
		if e.lessionLenght > 0 && e.stats.TotalTasks == e.lessionLenght {
			break
		}

		task, err := e.lession.NextTask()
		if err != nil {
			if errors.Is(err, ErrNoNextTask) {
				break
			}

			return e.stats, err
		}

		err = e.outputer.OutTarget(task.Target())
		if err != nil {
			return e.stats, err
		}

		err = e.tts.Say(task.Target())
		if err != nil {
			return e.stats, err
		}

		err = e.outputer.OutPrompt()
		if err != nil {
			return e.stats, err
		}

		answerFromUser, err := e.waiter.WaitAnswer()
		if err != nil {
			return e.stats, err
		}

		ok := strings.ToLower(answerFromUser) == strings.ToLower(task.Answer())

		if ok {
			e.stats.OkTaksks++
		}

		e.stats.TotalTasks++

		err = e.outputer.OutAnswer(task.Answer(), ok)
		if err != nil {
			return e.stats, err
		}

		err = e.tts.Say(task.Answer())
		if err != nil {
			return e.stats, err
		}
	}

	return e.stats, nil
}
