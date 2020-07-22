package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {
	// Open the file
	csvfile, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	reader := bufio.NewReader(os.Stdin)
	counter := 1
	score := 0

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Problem #%d: %s = ", counter, record[0])
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

	    if strings.Compare(record[1], text) == 0 {
			score += 1
		}

		counter += 1
	}

	fmt.Printf("Score: %d\n", score);
}

