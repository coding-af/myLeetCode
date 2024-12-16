package main

import (
	"fmt"
	"github.com/coding-af/myLeetCode/easy"
)

func main() {
	//check perfix of world
	out := easy.IsPrefixOfWord("i love eating burger", "burg")
	fmt.Println(out)
	//running sum array 1D
	input := []int{1, 2, 3, 4}
	sumOut := easy.RunningSum(input)
	fmt.Println(sumOut)
	//fizzbuzz
	fizzBuzzOut := easy.FizzBuzz(15)
	fmt.Println(fizzBuzzOut)
	//number of step
	numberOfStep := easy.NumberOfSteps(123)
	fmt.Println(numberOfStep)
}
