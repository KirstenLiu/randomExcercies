package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//parsing cmd: -h, -csv, -limit(Q2)

func readInput(quit chan string){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(strings.TrimSpace(input))
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	quit<-input
}

func main(){
	args := os.Args
	//fmt.Println("args: ", args[1])
	fileName := "problems.csv"
	timeOut := 30
	timeOutCh := make(chan string, 1)
	start := false

	shuffle := false
	for index, arg := range(args){
		//fmt.Println(index, len(args))
		if string(arg[0]) == "-" {
			arg = strings.Trim(arg, "-")
			switch arg {
				case "h":
					fmt.Println("-csv string\n\ta csv file in format of 'question, answer' (default problems.csv)")
					fmt.Println("-limit\n\tthe time limit for the quiz in seconds (default 30)")
				case "csv":
					if index+1 < len(args) {
						fileName = args[index+1]
					}
					start = true
				case "limit":
					//fmt.Println("This is for Q2")
					if index+1 < len(args) {
						timeOutInt, err := strconv.Atoi(args[index+1])
						if err != nil {
							log.Fatal(err)
						}
						timeOut = timeOutInt
					}
					//fmt.Println("TimeOut is", timeOut)
					start = true
				
				case "shuffle":
					shuffle = true
					start = true
					rand.Seed(time.Now().UnixNano())
				default:
					fmt.Println("Please enter a valid cmd. Cmd list: -h|-csv|-limit")
			}
		}
	}	
	correct := 0

	if start ==true {
		csvFile, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)
		csvLines, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		if shuffle == true {
			rand.Shuffle(len(csvLines), func(i, j int){csvLines[i], csvLines[j] = csvLines[j], csvLines[i]})
		}

		for _, line := range (csvLines){
			//fmt.Println(line[0])
			//fmt.Println(line[1])
			question := line[0]
			ans := line[1]

			fmt.Println(question)
			go readInput(timeOutCh)
			select {
				case res := <- timeOutCh:
					if res == ans {
						correct = correct + 1
					}
				case <- time.After(time.Duration(timeOut)*time.Second):
					fmt.Println("time out")
			}
		}
		fmt.Printf("your result is: %d correct in %d questions\n", correct, len(csvLines))
	}
}