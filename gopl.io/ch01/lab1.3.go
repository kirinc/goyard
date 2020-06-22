package main

import(
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start1 := time.Now()
	var s1, sep1 string
	for i := 1; i<len(os.Args); i++{
		s1 += sep1 + os.Args[i]
		sep1 = " "
	}
	fmt.Println(s1)
	seconds1 := time.Since(start1).Seconds()
	fmt.Printf("%.10f\n", seconds1)

	start2 := time.Now()
	s2, sep2 := "", ""
	for _, arg := range(os.Args[1:]){
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	seconds2 := time.Since(start2).Seconds()
	fmt.Printf("%.10f\n", seconds2)

	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:] , " "))
	seconds3 := time.Since(start3).Seconds()
	fmt.Printf("%.10f\n", seconds3)

}
