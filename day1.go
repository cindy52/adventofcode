package main

import (
	"strconv"
	"flag"
	"log"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "1122", "input string")
	flag.Parse()

	sum := 0
	l := len(input)
	for i, _ := range input {
		if i+1 < l && input[i] == input[i+1] {
			r, _:= strconv.Atoi(string(input[i]))
			sum += r
		}
	}
	if input[0] == input[l-1] {
		r, _:= strconv.Atoi(string(input[0]))
		sum += r
	}
	log.Println(sum)
}
