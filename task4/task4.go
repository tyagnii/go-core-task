package main

import "fmt"

func Exclusion(first, second []string) []string {
	var result []string
	var itsHere bool = false

	for _, v := range first {
		for _, w := range second {
			if v == w {
				itsHere = true
			}
		}
		if itsHere == false {
			result = append(result, v)
		}
		itsHere = false
	}
	return result
}

func main() {
	f := []string{"one", "two", "three", "four", "five", "six"}
	s := []string{"one", "two", "four", "five"}
	r := Exclusion(f, s)
	fmt.Println(r)
}
