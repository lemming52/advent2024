package main

import (
	"advent/solutions/d01hysteria"
	"advent/solutions/d02rednosereports"
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

    }
	return res
}
