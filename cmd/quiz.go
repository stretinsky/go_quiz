package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type question string
type answer string

type problem struct {
	question question
	answer   answer
}

type quiz struct {
	problems              []problem
	shuffle               bool
	countCorrectAnswers   int
	countIncorrectAnswers int
}

func (q *quiz) start() {
	if q.shuffle {
		q.doShuffle()
	}

	for i := 0; i < len(q.problems); i++ {
		var answer answer
		fmt.Println(q.problems[i].question)
		_, err := fmt.Scan(&answer)
		if err != nil {
			log.Fatalf("Unable to read answer")
		}
		q.checkAnswer(answer, q.problems[i])
	}
}

func (q *quiz) loadProblems(problems [][]string) {
	q.problems = []problem{}
	for _, i := range problems {
		q.problems = append(q.problems, problem{
			question: question(i[0]),
			answer:   answer(i[1]),
		})
	}
}

func (q *quiz) printResult() {
	fmt.Println("Количество правильных ответов: ", q.countCorrectAnswers)
	fmt.Println("Количество неправильных ответов: ", q.countIncorrectAnswers)
}

func (q *quiz) checkAnswer(answer answer, p problem) {
	if strings.ToLower(strings.TrimSpace(string(answer))) == strings.ToLower(strings.TrimSpace(string(p.answer))) {
		q.countCorrectAnswers++
	} else {
		q.countIncorrectAnswers++
	}
}

func (q *quiz) doShuffle() {
	for i := range q.problems {
		j := rand.Intn(i + 1)
		q.problems[i], q.problems[j] = q.problems[j], q.problems[i]
	}
}
