package main

import(
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//lab 1.3 + original
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	seconds := time.Since(start).Seconds()
	fmt.Printf("%.10f", seconds)

	//lab 1.1
	fmt.Println(strings.Join(os.Args[0:], " "))

	//lab 1.2

}
