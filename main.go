package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))

}
