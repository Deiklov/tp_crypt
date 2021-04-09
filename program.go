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

type resp struct {
	Name         string
	TicketNumber int
}

func main() {

	var fileName = flag.String("file", "program", "The name of input files")
	var numBilets = flag.Int("numbilets", 20, "Count of tickets for students")
	var parameter = flag.Int("parameter", 42, "Parameter for change seed")

	// parse flag's options
	flag.Parse()

	f, err := os.OpenFile(*fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	arrFIOs := []string{}
	for sc.Scan() {
		currFIO := sc.Text()
		arrFIOs = append(arrFIOs, currFIO)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	tickets, err := genTickets(arrFIOs, *parameter, *numBilets)
	if err != nil {
		log.Fatal(err)
	}
	//имя это index
	for _, v := range tickets {
		fmt.Printf("%s: %d \n", v.Name, v.TicketNumber)
	}

}
func genTickets(studFIOs []string, seed int, numTickets int) ([]resp, error) {
	response := make([]resp, len(studFIOs))

	sort.Strings(studFIOs)
	rand.Seed(int64(seed))

	//генерит массив рандомных чисел из отрезка [1, N] длинной равной числу студентов
	ticketsArray := func() []int {
		arr := make([]int, len(studFIOs))
		for i := 0; i < len(studFIOs); i++ {
			arr[i] = rand.Intn(numTickets) + 1
		}
		return arr
	}()

	//каждому студенту ставим в соответствие его номер билета
	for i, v := range ticketsArray {
		response[i].Name = studFIOs[i]
		response[i].TicketNumber = v
	}

	return response, nil
}
