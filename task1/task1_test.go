package main

import "testing"

func TestIdentType(t *testing.T) {
	var a int
	if IdentType(a) != "int" {
		t.Error("wrong type")
	}
}

func TestMakeStringFromAny(t *testing.T){
	var f float64 = 4.23
	var b bool = true
	if MakeStringFromAny(f,b) != "4.23true" {
		t.Error("incorrect result string")
	}
}

func TestStrToRunes(t *testing.T) {
	var s string = "some string"
	r := []rune(s)
	res := StrToRunes(s)
	for res
	}
}
