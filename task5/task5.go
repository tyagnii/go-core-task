package main

import "fmt"

func Intersection(first, second []int) ([]int, bool) {
	var flag bool = false

	result := make([]int, 0)
	tempMap := make(map[int]bool)

	for _, v := range first {
		for _, w := range second {
			if v == w {
				tempMap[w] = true
			}
		}
	}

	if len(tempMap) == 0 {
		return result, flag
	} else {

		flag = true
		for v, _ := range tempMap {
			result = append(result, v)
		}
		return result, flag
	}
}

func main() {
	f := []int{65, 3, 58, 678, 64}
	s := []int{64, 2, 3, 43}
	r, _ := Intersection(f, s)
	fmt.Println(r)
}
