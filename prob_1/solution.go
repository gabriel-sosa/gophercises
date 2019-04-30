package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

var pathToQuestions string
var timePerAnswer int

func main() {
	parseFlags()
	file, err := getQuestions()
	if err != nil {
		fmt.Println("error")
	} else {
		score := answerQestions(file)
		fmt.Println("your  score is", score)
	}
}

func parseFlags() {
	flag.StringVar(&pathToQuestions, "path", "./problems.csv", "path to the csv file")
	flag.IntVar(&timePerAnswer, "time", 30, "time available for each answer")
	flag.Parse()
}

func getQuestions() ([]byte, error) {
	return ioutil.ReadFile(pathToQuestions)
}

func answerQestions(file []byte) (res int) {
	res = 0
	r := csv.NewReader(strings.NewReader(string(file)))
	for {
		record, err := r.Read()
		answer := make(chan string)
		if err == io.EOF {
			return
		}
		go func() {
			var input string
			fmt.Println(record[0])
			fmt.Scanln(&input)
			answer <- input
		}()
		select {
		case input := <-answer:
			if record[1] == input {
				res++
			}
		case <-time.After(time.Duration(timePerAnswer) * time.Second):
			return
		}
	}
}
