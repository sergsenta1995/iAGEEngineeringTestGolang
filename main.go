package main

import (
	"fmt"
	"strconv"
	"errors"
	"bufio"
	"os"
)

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

var triads = [][]string{
	// падежи
	{"",         "",          ""},
	{"тысяча ",  "тысячи ",   "тысяч "},
	{"миллион ", "миллиона ", "миллионов "}}

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
	"пятьдесят ",
	"шестьдесят ",
	"семьдесят ",
	"восемьдесят ",
	"девяносто "}

var thousands = []string{
	"сто ",
	"двести ",
	"триста ",
	"четыреста ",
	"пятьсот ",
	"шестьсот ",
	"семьсот ",
	"восемьсот ",
	"девятьсот "}

// Особы индекс для тысячи, т. к. тысячи, заканчивающиеся на '1' имеют женски род
const thousand = 1

func IntToNumeric(number int) string {
	var triadNumeric string
	var result string

	if number == 0 {
		return "ноль"
	}

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
			// Только тысяча, заканчивающаяся на '1' может иметь женский род
			if triadNumber == thousand {
				triadNumeric += "одна "
				triadNumeric += triads[triadNumber][0]
			} else {
				triadNumeric += ones[triad-1]
				triadNumeric += triads[triadNumber][0]
			}
		} else {
			if triad >= 2 && triad <= 4 {
				triadNumeric += ones[triad-1]
				triadNumeric += triads[triadNumber][1]
			} else {
				if triad >= 5 && triad <= 9 {
					triadNumeric += ones[triad-1]
					triadNumeric += triads[triadNumber][2]
				}
			}
		}
		result = triadNumeric + result
		triadNumeric = ""
	}
	// Убрать лишний пробельный символ
	return result[0:len(result)-1]
}

func CheckNumber(stringNumber string) (intNumber int, err error) {
	intNumber, err = strconv.Atoi(stringNumber)
	if err != nil {
		return 0, err
	}
	if intNumber < 0 {
		return 0, errors.New("Ошибка: число меньше нуля")
	}
	if intNumber > 1000000 {
		return 0, errors.New("Ошибка: число больше миллиона")
	}
	return intNumber, err
}

func main ()  {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println("Введите число от 0 до 1000000: ")
	stringNumber, _:= reader.ReadString('\n')
	// Убрать символ конца строки - без него не преобразуется в число
	stringNumber = stringNumber[0:len(stringNumber)-1]
	
	intNumber, err := CheckNumber(stringNumber)
	if err == nil {
		numeric := IntToNumeric(intNumber)
		fmt.Println(numeric)
	} else {
		fmt.Println(err)
	}
}
