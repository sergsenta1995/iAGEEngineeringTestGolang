package main

import "fmt"

var ones = []string {
	"один ",
	"два ",
	"три ",
	"четыре ",
	"пять ",
	"шесть ",
	"семь ",
	"восемь ",
	"девять "}

var triads = []string{
	"",
	"тысяча ",
	"миллион "}

var onesTens = []string{
	"одиннадцать ",
	"двенадцать ",
	"тринадцать ",
	"четырнадцать ",
	"пятнадцать ",
	"шестнадцать ",
	"семнадцать ",
	"воесемнадцать ",
	"девятнадцать "}

var tens = []string{
	"десять ",
	"двадцать ",
	"тридцать ",
	"сорок ",
	"пятдесят ",
	"шестдесят ",
	"семдесят ",
	"восемдесят ",
	"девяносто "}

var thousands = []string{
	"сто ",
	"двести ",
	"триста ",
	"четыреста ",
	"пятсот ",
	"шестьсот ",
	"семсот ",
	"восемсот ",
	"девятсот "}

func IntToNumeric(number int) string {
	var triadNumeric string
	var result string
	for triadNumber := 0; number >= 1; triadNumber++ {
		triad := number % 1000
		number /= 1000
		if triad >= 100 {
			triadNumeric += thousands[triad/100-1]
			triad %= 100
		}
		if triad >= 20 || triad == 10 {
			triadNumeric += tens[triad/10-1]
			triad %= 10
		}
		if triad >= 11 {
			triadNumeric += onesTens[triad/10-1]
		}
		if triad == 1 {
			triadNumeric += ones[triad-1]
			triadNumeric += triads[triadNumber]
		}
		result = triadNumeric + result
		triadNumeric = ""
	}
	return result[0:len(result)-1]
}

func main ()  {
	fmt.Println(IntToNumeric(1000))
}
