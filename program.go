package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
)

func main() {

	// to save the 'name' option value
	var fileName = flag.String("file", "program", "The name of input files")
	var numBilets = flag.Int("numbilets", 20, "Count of tickets for students")
	var randSeed = flag.Int("parameter", 42, "Count of tickets for students")

	// parse flag's options
	flag.Parse()

	f, err := os.OpenFile(*fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	arrFIO := []string{}
	for sc.Scan() {
		currStudent := sc.Text()
		arrFIO = append(arrFIO, currStudent)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	ticketsMAP, err := genTicket(arrFIO, *randSeed, *numBilets)
	if err != nil {
		log.Fatal(err)
	}
	//имя это index
	for i, v := range ticketsMAP {
		fmt.Printf("%s: %d \n", i, v)
	}

}
func genTicket(studFIOs []string, seed int, numTickets int) (map[string]int, error) {
	//все имена будут под одной маской
	response := make(map[string]int, len(studFIOs))

	sort.Strings(studFIOs)
	rand.Seed(int64(seed))

	ticketsArray := func() []int {
		arr := make([]int, len(studFIOs))
		for i := 0; i < len(studFIOs); i++ {
			arr[i] = rand.Intn(numTickets)+1
		}
		return arr
	}()

	for i, v := range ticketsArray {
		response[studFIOs[i]] = v
	}
	//создали сид для рандома
	return response, nil
}
