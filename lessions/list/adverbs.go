package list

import (
	"maps"
	"math/rand/v2"
	"slices"
	"sort"

	"github.com/name212/en-lang-learner/pkg"
	"github.com/name212/en-lang-learner/pkg/lession"
)

type AdverbsLession struct {
	*lession.BiderictionalRandomLession
}

func randomElement(m map[string]string) string {
	keys := slices.Collect(maps.Keys(m))
	keys = sort.StringSlice(keys)
	indx := rand.IntN(len(keys))
	k := keys[indx]
	return m[k]
}

var (
	professions = map[string]string{
		"teacher":    "учитель",
		"driver":     "водитель",
		"painter":    "художник",
		"model":      "модель",
		"writer":     "писатель",
		"backer":     "пекарь",
		"blogger":    "блоггер",
		"head":       "глава",
		"lawyer":     "судья",
		"loader":     "грузчик",
		"trainer":    "тренер",
		"waiter":     "официант",
		"plumber":    "сантехник",
		"programmer": "программист",
	}
)

func NewAdverbsLession() pkg.Lession {
	return lession.NewBiderictionalWithExceptionsRandomLession([]pkg.Task{
		pkg.NewBaseTask("I", "я"),
		pkg.NewBaseTask("he", "он"),
		pkg.NewBaseTask("she", "она"),
		pkg.NewBaseTask("it", "оно"),
		pkg.NewBaseTask("we", "мы"),
		pkg.NewBaseTask("they", "они"),

		pkg.NewBaseTask("me", "мне"),
		pkg.NewBaseTask("him", "ему"),
		pkg.NewBaseTask("us", "нам"),
		pkg.NewBaseTask("them", "им"),

		pkg.NewBaseTask("my", "мой"),
		pkg.NewBaseTask("his", "его"),
		pkg.NewBaseTask("its", "его, её"),
		pkg.NewBaseTask("our", "наше"),
		pkg.NewBaseTask("their", "их"),
	}, []pkg.Task{
		pkg.NewBaseTask("ты", "you"),
		pkg.NewBaseTask("Вы (множ.)", "you"),

		pkg.NewBaseTask("тебе", "you"),
		pkg.NewBaseTask("Вам", "you"),
		pkg.NewBaseTask("ему, ей (ср. род)", "it"),
		pkg.NewBaseTask("ей", "her"),

		pkg.NewBaseTask("её", "her"),
		pkg.NewBaseTask("твой", "your"),
		pkg.NewBaseTask("ваше", "your"),

		pkg.NewBaseTask("mine", "мой"),
		pkg.NewBaseTask("hers", "её"),
		pkg.NewBaseTask("yours", "твой, ваш"),
		pkg.NewBaseTask("ours", "наш"),
		pkg.NewBaseTask("theirs", "их"),

		pkg.NewBaseTask("myself", "I"),
		pkg.NewBaseTask("yourself", "you"),
		pkg.NewBaseTask("himself", "he"),
		pkg.NewBaseTask("herself", "she"),
		pkg.NewBaseTask("itself", "it"),
		pkg.NewBaseTask("ourselves", "we"),
		pkg.NewBaseTask("yourselves", "you"),
		pkg.NewBaseTask("themselves", "they"),
	})
}
