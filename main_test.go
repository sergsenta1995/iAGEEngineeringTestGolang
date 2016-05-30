package main

import (
	"testing"
)

type testpair struct {
	value int
	numeric string
}

var tests = []testpair {
	{1, "один"},
	{10, "десять"},
	{100, "сто"},
	{1000, "одна тысяча"},
	{123456, "сто двадцать три тысячи четыресто пятдесят шесть"}}

func TestIntToNumeric(t *testing.T) {
	for _, pair := range tests {
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