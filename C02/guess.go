package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	gen := rand.New(rand.NewSource(100))
	target := gen.Float64()

	fmt.Println("Please input your guess:")
	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var status string
	input = strings.TrimSpace(input)
	guess, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	if guess < target {
		status = "Oops, your guess was LOW!"
	} else if guess > target {
		status = "Oops, your guess was HIGH!"
	} else {
		status = "Bingo, you did it!!!"
	}
	fmt.Println(status)
}
