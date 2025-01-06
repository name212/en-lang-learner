package tts

import "github.com/name212/en-lang-learner/pkg/cli"

type EspeakNGSpeetcher struct{}

func NewEspeakNGSpeetcher() *EspeakNGSpeetcher {
	return &EspeakNGSpeetcher{}
}

func (s EspeakNGSpeetcher) Say(text string) error {
	_, err := cli.NewExecutor().Exec("espeak-ng", text)
	return err
}
