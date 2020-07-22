package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
	"strings"
	"time"
	"flag"
)


var score = 0

func main() {
    c1 := make(chan string, 1)
	timeout := flag.Int("timeout", 5, "timeout of quiz")
	fileName := flag.String("csv", "quiz.csv", "quiz file")
	flag.Parse()

    // Run your long running function in it's own goroutine and pass back it's
    // response into our channel.
    go func() {
        quiz(*fileName)
    }()

    // Listen on our channel AND a timeout channel - which ever happens first.
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Duration(*timeout) * time.Second):
		csvfile, _ := os.Open(*fileName)
		r := csv.NewReader(csvfile)
		lines, _ := r.ReadAll()

        fmt.Println("\n\nout of time :(")
		fmt.Printf("Scored %d out of %d\n", score, len(lines));
    }
}

func quiz(fileName string){
	// Open the file
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	reader := bufio.NewReader(os.Stdin)

	counter := 1
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

