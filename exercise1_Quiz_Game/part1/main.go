package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	problems, err := readFile(*fileName)
	if err != nil {
		panic(err)
	}

	correct := 0
	for i, problem := range problems {
		ps := strings.Split(problem, ",")
		var answer string
		fmt.Printf("Problem #%d: %s = ", i+1, ps[0])
		fmt.Scanf("%s \n", &answer)
		if answer == ps[1] {
			correct++
		}
	}
	fmt.Printf("Total correct answer is: %d\n", correct)
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
