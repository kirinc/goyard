package main

import(
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	seconds :=time.Since(start).Seconds()
	fmt.Printf("%.10f", seconds)

	//lab 1.2
	for index, arg := range os.Args[1:]{
		fmt.Println(index, arg)
	}
}
