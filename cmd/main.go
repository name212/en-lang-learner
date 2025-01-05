package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/name212/en-lang-learner/lessions"
	"github.com/name212/en-lang-learner/pkg"
	"github.com/name212/en-lang-learner/pkg/cli"
	"github.com/name212/en-lang-learner/pkg/tts"
)

func printLessionsAndExit() {
	fmt.Printf("Lessions list:\n")
	for _, l := range lessions.List {
		fmt.Printf("\t%-4s - %s\n", l.Name, l.Description)
	}

	os.Exit(0)
}

func printEerrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func printStatistic(s *pkg.Stats) {
	percent := (float64(s.OkTaksks) / float64(s.TotalTasks)) * 100
	fmt.Printf("Total: %d\n", s.TotalTasks)
	fmt.Printf("Ok: %d(%.2f%%)\n", s.OkTaksks, percent)
}

func main() {
	listLessionsFlag := flag.Bool("show-lessions-list", false, "Show available lessions list")
	lessionName := flag.String("name", "", "Select lession for start")
	lessionLenght := flag.Int("lession-len", 0, "Limit number of tasks in the lession. 0 - no limit")

	flag.Parse()

	if *listLessionsFlag || *lessionName == "" {
		printLessionsAndExit()
		return
	}

	lession, err := lessions.FindLessionByName(*lessionName)
	if err != nil {
		printEerrorAndExit(err)
		return
	}

	tts := tts.NewEspeakNGSpeetcher()
	outputer := cli.NewOutputer()
	waiter := cli.NewAnswerWaiter()

	e := pkg.NewEngine(lession.Lession(), tts, outputer, waiter).
		WithLessionLenght(*lessionLenght)

	stats, err := e.Run()
	if err != nil {
		printEerrorAndExit(err)
	}

	printStatistic(stats)

}
