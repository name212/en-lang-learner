package lessions

import "github.com/name212/en-lang-learner/pkg"

type testLession struct {
	*linearLession
}

func newTestLession() pkg.Lession {
	return &testLession{
		linearLession: newLinearLession([]pkg.Task{
			pkg.NewBaseTask("Hello", "Привет"),
			pkg.NewBaseTask("World", "Мир"),
			pkg.NewBaseTask("От", "From"),
			pkg.NewBaseTask("Меня", "Me"),
		}),
	}
}
