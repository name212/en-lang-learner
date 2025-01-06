package lessions

import (
	"fmt"

	lst "github.com/name212/en-lang-learner/lessions/list"
	"github.com/name212/en-lang-learner/pkg"
)

type Item struct {
	Description string
	Lession     func() pkg.Lession
	Name        string
}

var List = []Item{
	Item{
		Name:        "adverbs",
		Description: "Adverbs repeating",
		Lession:     lst.NewAdverbsLession,
	},
	Item{
		Name:        "test-linear",
		Description: "Lession for test purposes",
		Lession:     lst.NewTestLinearLession,
	},
	Item{
		Name:        "test-bidirectionalrandom",
		Description: "Lession for test purposes",
		Lession:     lst.NewTestBidirectionalRandomLession,
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
