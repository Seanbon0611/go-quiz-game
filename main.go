package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type questionsAndAnswers struct {
	Question string
	Answer   string
}

func main() {
	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader, _ := csv.NewReader(file).ReadAll()
	fmt.Println(reader)

	for _, line := range reader {
		qAndA := questionsAndAnswers{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(qAndA.Question + " " + qAndA.Answer)
	}
}
