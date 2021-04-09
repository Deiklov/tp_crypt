package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {

	// to save the 'name' option value
	var fileName = flag.String("file", "program", "The name of input files")
	var numBilets = flag.Int64("numbilets", 20, "Count of tickets for students")
	var randSeed = flag.Int64("parameter", 42, "Count of tickets for students")

	// parse flag's options
	flag.Parse()

	f, err := os.OpenFile(*fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		currStudent := sc.Text()
		ticketNumber, err := genTicket(currStudent, *randSeed, *numBilets)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("%s: %d", currStudent, ticketNumber)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

}
func genTicket(studFIO string, seed int64, numTickets int64) (int64, error) {
	//все имена будут под одной маской
	hashName := sha256.Sum256([]byte(studFIO))
	//создали сид для рандома
	rand.Seed(seed)
	return 0, nil
}
