package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// generateSlice generates slice of 10 ints from 0 to 100
func generateSlice() []int {
	var originalSlice = make([]int, 10)
	for i := 0; i < 10; i++ {
		originalSlice[i] = rand.Intn(100)
	}
	return originalSlice
}

// sliceExample removes even numbers from slice
func sliceExample(originalSlice []int) []int {
	var result []int

	for _, v := range originalSlice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// addElements is custom Append func for slices
func addElements(originalSlice []int, elem int) []int {
	if len(originalSlice) >= cap(originalSlice) {
		newSlice := make([]int, len(originalSlice)+1, 2*len(originalSlice)+1)
		for i, v := range originalSlice {
			newSlice[i] = v
		}
		newSlice[len(originalSlice)] = elem
		return newSlice

	} else {
		originalSlice[len(originalSlice)+1] = elem
		return originalSlice
	}

}

// copySlice is custom copy func
func copySlice(originalSlice []int) []int {
	var copyOfSlice = make([]int, len(originalSlice))
	for i := 0; i < len(originalSlice); i++ {
		copyOfSlice[i] = originalSlice[i]
	}

	return copyOfSlice
}

// removeElement removes element from slice by index
func removeElement(originalSlice []int, index int) ([]int, error) {
	var result []int

	if index > len(originalSlice) {
		return originalSlice, errors.New("index out of range")
	}

	result = append(originalSlice[:index], originalSlice[index+1:]...)
	return result, nil
}

func main() {
	oSlice := generateSlice()
	fmt.Println(oSlice)

	slice1 := addElements(oSlice, 50)
	fmt.Println(slice1)

	slice2 := sliceExample(slice1)
	fmt.Println(slice2)

	copyOfSlice := copySlice(oSlice)
	fmt.Println(copyOfSlice)

	rSlice, err := removeElement(copyOfSlice, 200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rSlice)

}
