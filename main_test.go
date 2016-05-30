package main

import (
	"testing"
)

type testPairForIntToNumeric struct {
	value int
	numeric string
}

var testsForIntToNumeric = []testPairForIntToNumeric{
	{1, "один"},
	{10, "десять"},
	{100, "сто"},
	{1000, "одна тысяча"},
	{123456, "сто двадцать три тысячи четыреста пятдесят шесть"}}

func TestIntToNumeric(t *testing.T) {
	for _, pair := range testsForIntToNumeric {
		result := IntToNumeric(pair.value)
		if result != pair.numeric {
			t.Error(
				"Для значения <", pair.value,">;",
				"допустиый результат <", pair.numeric,">;",
				"Получено <", result,">",
			)
		}
	}
}

type testPairForCheckNumber struct {
	stringNumber string
	intNumber int
}

var testsForCheckNumber = []testPairForCheckNumber{
	{"1", 1},
	{"-9", 0},
	{"123,", 0},
	{"125698777", 0}}

func TestCheckNumber(t *testing.T) {
	for _, pair := range testsForCheckNumber {
		result, _ := CheckNumber(pair.stringNumber)
		if result != pair.intNumber {
			t.Error(
				"Для значения <", pair.stringNumber,">;",
				"допустиый результат <", pair.intNumber,">;",
				"Получено <", result,">",
			)
		}
	}
}