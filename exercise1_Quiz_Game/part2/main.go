package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	fmt.Println(*limit)
	problems, err := readFile(*fileName)
	if err != nil {
		panic(err)
	}
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0

executionLoop:
	for i, problem := range problems {
		ps := strings.Split(problem, ",")

		fmt.Printf("Problem #%d: %s = ", i+1, ps[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break executionLoop
		case answer := <-answerCh:
			if answer == ps[1] {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func readFile(fileName string) ([]string, error) {
	result := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || len(line) == 0 {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		result = append(result, line)
	}
	return result, nil
}
