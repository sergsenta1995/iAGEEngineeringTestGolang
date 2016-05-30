package main

import (
	"testing"
)

func TestIntToNumeric(t *testing.T) {
	if IntToNumeric(1) != "один" {
		t.Error("Число состоит из одного разряда.")
		t.Error("Ошибка в определении первого разряда")
	}
}