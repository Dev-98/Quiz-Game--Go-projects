package main

import (
	"encoding/csv"
	f "flag"
	fm "fmt"
	// "log"
	"os"
	"time"
)

// func Querygetter(path string) (lines []string, err error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal("Could not open file: ", path)
// 	}
// 	defer file.Close()
// 	data, err := csv.NewReader(file).ReadAll()
// 	if err != nil {
// 		log.Fatal("Could not read file data ")
// 	}
// 	return locations, nil
// }

func main() {
	filename := f.String("csv", "problems.csv", "a CSV file containing problems in the format of 'question,answer'")
	timeduration := f.Int("timer", 30, "How long you get to answer all the questions")

	f.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		exit(fm.Sprintf("Can't open file: %s\n", *filename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll()

	if err != nil {
		exit("Could not read file data ")
	}

	count := 0
	timer := time.NewTimer(time.Duration(*timeduration) * time.Second)

	for i, val := range data {
		fm.Println("YOU HAVE 30 sec to ANSWER | HURRY!! \n\nQuestion #", i+1, ": ", val[0])
		fm.Print("Your answer: ")

		answerCh := make(chan string)
		go func() {
			var input string
			fm.Scanln(&input)
			answerCh <- input
		}()

		select {
		case <-timer.C:
			fm.Print("\n=========\nTimes up!!")
			fm.Printf("\nYou scored %d out of %d\n", count, len(data))
			return

		case answer := <-answerCh:
			if answer == val[1] {
				count++
			}
		}

	}
	fm.Printf("\nYou scored %d out of %d\n", count, len(data))
}

func exit(msg string) {
	fm.Println(msg)
	os.Exit(1)
}
