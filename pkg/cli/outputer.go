package cli

import "fmt"

type Outputer struct{}

func NewOutputer() *Outputer {
	return &Outputer{}
}

func (o Outputer) OutTarget(t string) error {
	fmt.Printf("Translate: %s\n", t)
	return nil
}

func (o Outputer) OutAnswer(a string, isTrue bool) error {
	doneFlag := "✓"
	if !isTrue {
		doneFlag = "❌"
	}

	fmt.Printf("%s Answer: %s\n", doneFlag, a)

	return nil
}

func (o Outputer) OutPrompt() error {
	fmt.Printf("-> ")
	return nil
}
