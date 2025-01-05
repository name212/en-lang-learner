package lessions

import (
	"fmt"

	"github.com/name212/en-lang-learner/pkg"
)

type Item struct {
	Description string
	Lession     func() pkg.Lession
	Name        string
}

var List = []Item{
	Item{
		Name:        "test",
		Description: "Lession for test purposes",
		Lession:     newTestLession,
	},
}

func FindLessionByName(name string) (Item, error) {
	for _, l := range List {
		if l.Name == name {
			return l, nil
		}
	}

	return Item{}, fmt.Errorf("Lession %s not found", name)
}
