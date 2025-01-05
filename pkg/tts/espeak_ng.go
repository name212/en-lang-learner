package tts

type EspeakNGSpeetcher struct{}

func (s EspeakNGSpeetcher) Say(text string) error {
	_, err := newRunner().exec("espeak-ng", text)
	return err
}
