package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

var numDecimal int = 42       // Десятичная система
var numOctal int = 052        // Восьмеричная система
var numHexadecimal int = 0x2A // Шестнадцатиричная система
var pi float64 = 3.14
var e float32 = 2.72              // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

func IdentType(n any) string {
	return reflect.TypeOf(n).String()

}

func MakeStringFromAny(a ...any) string {
	var str string
	for _, i := range a {
		str += fmt.Sprint(i)
	}
	return str
}

func StrToRunes(s string) []rune {
	return []rune(s)

}

func SHA256(runes []rune, salt string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(string(runes)))
	if err != nil {
		fmt.Println("Hash error:", err)
	}
	return string(hash.Sum([]byte(salt)))
}

func main() {
	fmt.Println("Task #1")

	// Identify type of var
	IdentType(numDecimal)
	IdentType(numOctal)
	IdentType(numHexadecimal)
	IdentType(pi)
	IdentType(name)
	IdentType(isActive)
	IdentType(complexNum)

	// Convert batch of vars to string
	str := MakeStringFromAny(numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		e,
		name,
		isActive,
		complexNum)
	fmt.Println("Any to one string:", str)

	// Convert string to runes
	runes := StrToRunes(str)
	fmt.Println("String to slice of runes:", runes)

	// Create hash
	salt := "go-2024"
	fmt.Printf("%x", SHA256(runes, salt))

}
