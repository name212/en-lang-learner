package lessions

import (
	"github.com/name212/en-lang-learner/pkg"
	"github.com/name212/en-lang-learner/pkg/lession"
)

type TestLinearLession struct {
	*lession.LinearLession
}

func NewTestLinearLession() pkg.Lession {
	return &TestLinearLession{
		LinearLession: lession.NewLinearLession([]pkg.Task{
			pkg.NewBaseTask("Hello", "Привет"),
			pkg.NewBaseTask("World", "Мир"),
			pkg.NewBaseTask("От", "From"),
			pkg.NewBaseTask("Меня", "Me"),
		}),
	}
}

type TestBidirectionalRandomLession struct {
	*lession.BiderictionalRandomLession
}

func NewTestBidirectionalRandomLession() pkg.Lession {
	return &TestBidirectionalRandomLession{
		BiderictionalRandomLession: lession.NewBiderictionalRandomLession([]pkg.Task{
			pkg.NewBaseTask("Hello", "Привет"),
			pkg.NewBaseTask("World", "Мир"),
		}),
	}
}
