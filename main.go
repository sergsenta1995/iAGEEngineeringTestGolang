package main

import "fmt"

var ones = []string {
	"один",
	"два",
	"три",
	"четыре",
	"пять",
	"шесть",
	"семь",
	"восемь",
	"девять",
}

func IntToNumeric(number int) string {
	return ones[number-1]
}

func main ()  {
	fmt.Println("!!!")
}
