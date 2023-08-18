package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filePath := flag.String("file", "problems.csv", "CSV file containing quiz questions and answers")
	timeLimit := flag.Int("timer", 30, "Time limit for each question in seconds")
	shuffle := flag.Bool("shuffle", false, "Shuffle the quiz order")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if *shuffle {
		shuffleQuestions(lines)
	}

	correct := 0
	total := len(lines)

	fmt.Println("Press Enter to start the quiz...")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	defer timer.Stop()

	for i, line := range lines {
		question := line[0]
		answer := strings.TrimSpace(line[1])

		var userAnswer string
		fmt.Printf("Question %d: %s = ", i+1, question)

		answerCh := make(chan string)

		go func() {
			fmt.Scanln(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			return
		case userAnswer = <-answerCh:
		}

		if userAnswer == answer {
			correct++
		}
	}

	fmt.Printf("\nYou got %d out of %d questions correct.\n", correct, total)
}

func shuffleQuestions(lines [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
}
