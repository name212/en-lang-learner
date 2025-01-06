package tts

import "github.com/name212/en-lang-learner/pkg/cli"

type SpdSayNGSpeetcher struct{}

func NewSpdSayNGSpeetcher() *SpdSayNGSpeetcher {
	return &SpdSayNGSpeetcher{}
}

func (s SpdSayNGSpeetcher) Say(text string) error {
	_, err := cli.NewExecutor().Exec("spd-say", "-w", text)
	return err
}
