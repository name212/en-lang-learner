package list

import (
	"github.com/name212/en-lang-learner/pkg"
	"github.com/name212/en-lang-learner/pkg/lession"
)

type AdverbsLession struct {
	*lession.BiderictionalRandomLession
}

func NewAdverbsLession() pkg.Lession {
	return lession.NewBiderictionalWithExceptionsRandomLession([]pkg.Task{
		pkg.NewBaseTask("I", "я"),
		pkg.NewBaseTask("you", "ты"),
		pkg.NewBaseTask("he", "он"),
		pkg.NewBaseTask("she", "она"),
		pkg.NewBaseTask("it", "оно"),
		pkg.NewBaseTask("we", "мы"),
		pkg.NewBaseTask("they", "они"),

		pkg.NewBaseTask("me", "мне"),
		pkg.NewBaseTask("him", "ему"),
		pkg.NewBaseTask("her", "ей"),
		pkg.NewBaseTask("us", "нам"),
		pkg.NewBaseTask("them", "им"),

		pkg.NewBaseTask("my", "им"),
		pkg.NewBaseTask("his", "им"),
		pkg.NewBaseTask("its", "им"),
		pkg.NewBaseTask("our", "им"),
		pkg.NewBaseTask("their", "им"),
	}, []pkg.Task{
		pkg.NewBaseTask("Вы (множ.)", "you"),

		pkg.NewBaseTask("тебе", "you"),
		pkg.NewBaseTask("Вам", "you"),
		pkg.NewBaseTask("ему, ей (ср. род)", "you"),

		pkg.NewBaseTask("ему, ей (ср. род)", "you"),
	})
}
