package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	//fmt.Println(target)

	read := bufio.NewReader(os.Stdin)
	success := false

	//allow user 10 times to try
	for count := 0; count < 10; count++ {
		fmt.Printf("You have %d times to try", 10-count)
		fmt.Println("\nPlease input your guess:")

		input, err := read.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		var status string
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			status = "Oops, your guess was LOW!"
		} else if guess > target {
			status = "Oops, your guess was HIGH!"
		} else {
			success = true
			fmt.Println("Bingo, you did it!!!")
			break
		}
		fmt.Println(status)
	}
	if !success {
		fmt.Println("Sorry, you did not guess my number. it was: ", target)
	}
}
