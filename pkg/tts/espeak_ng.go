package tts

type EspeakNGSpeetcher struct{}

func NewEspeakNGSpeetcher() *EspeakNGSpeetcher {
	return &EspeakNGSpeetcher{}
}

func (s EspeakNGSpeetcher) Say(text string) error {
	_, err := newRunner().exec("espeak-ng", text)
	return err
}
