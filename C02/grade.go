package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	if grade == 100 {
		fmt.Println("Perfect score")
	} else if grade >= 60 && grade < 100 {
		fmt.Println("Passed")
	} else if grade >= 0 && grade < 60 {
		fmt.Println("failed")
	} else {
		fmt.Println("Please input valid grade!!!")
	}

}
