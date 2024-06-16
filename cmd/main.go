package main

import (
	"flag"
)

func main() {
	problemsFilePath := flag.String("f", "problems.csv", "filepath")
	shuffle := flag.Bool("s", false, "shuffle?")
	flag.Parse()

	problems := readCsvFile(*problemsFilePath)

	quiz := quiz{
		shuffle: *shuffle,
	}
	quiz.loadProblems(problems)
	quiz.start()
	quiz.printResult()
}
