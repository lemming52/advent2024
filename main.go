package main

import (
	"advent/solutions/d01hysteria"
	"advent/solutions/d02rednosereports"
	"advent/solutions/d03mullitover"
	"advent/solutions/d04ceressearch"
	"advent/solutions/d05printqueue"
	"flag"
	"fmt"
	"time"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "trebuchet", "name or number of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"hysteria",
		"rednosereports",
		"mullitover",
		"ceressearch",
		"printqueue",
	}
	if *all {
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {
			s := RunChallenge(c)
			current := time.Now()
			fmt.Println(s, " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
			previous = current
		}
	} else {
		fmt.Println(RunChallenge(challenge))
	}
}

func RunChallenge(challenge string) string {
	var res string
	switch challenge {
	case "hysteria", "1":
        input := "inputs/d01hysteria.txt"
		A, B := d01hysteria.Run(input)
		res = fmt.Sprintf("hysteria Results A: %s B: %s", A, B)
	case "rednosereports", "2":
        input := "inputs/d02rednosereports.txt"
		A, B := d02rednosereports.Run(input)
		res = fmt.Sprintf("rednosereports Results A: %s B: %s", A, B)
	case "mullitover", "3":
        input := "inputs/d03mullitover.txt"
		A, B := d03mullitover.Run(input)
		res = fmt.Sprintf("mullitover Results A: %s B: %s", A, B)
	case "ceressearch", "4":
        input := "inputs/d04ceressearch.txt"
		A, B := d04ceressearch.Run(input)
		res = fmt.Sprintf("ceressearch Results A: %s B: %s", A, B)
	case "printqueue", "5":
        input := "inputs/d05printqueue.txt"
		A, B := d05printqueue.Run(input)
		res = fmt.Sprintf("printqueue Results A: %s B: %s", A, B)

    }
	return res
}
