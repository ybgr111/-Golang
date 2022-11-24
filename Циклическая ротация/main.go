package main

import (
	"fmt"
	"os"
	"strconv"
)

func Solution(A []int, K int) []int {

	if K < 0 || len(A) == 0 {
		return A
	}

	rs := len(A) - K%len(A)

	A = append(A[rs:], A[:rs]...)

	return A
}

func main() {

	var Str string
	var K int

	fmt.Print("Enter the size of the array:")

	fmt.Fscan(os.Stdin, &Str)

	N, err := strconv.Atoi(Str)
	if err != nil {
		fmt.Print("Invalid type")
		return
	}

	if N < 0 || N > 100 {
		fmt.Print("Invalid number")
		return
	}

	var A = make([]int, N)

	for i := 0; i < N; i++ {

		fmt.Printf("Enter %dth element: ", i)
		fmt.Fscan(os.Stdin, &Str)

		A[i], err = strconv.Atoi(Str)
		if err != nil {
			fmt.Print("Invalid type")
			return
		}

		if A[i] < -1000 || A[i] > 1000 {
			fmt.Print("Invalid number")
			return
		}
	}

	fmt.Print("Enter the number of shifts:")

	fmt.Fscan(os.Stdin, &Str)

	K, err = strconv.Atoi(Str)

	if err != nil {
		fmt.Print("Invalid type")
		return
	}

	if K < 0 || K > 100 {
		fmt.Print("Invalid number")
		return
	}

	fmt.Printf("slice %v\n", A)

	A = Solution(A, K)

	fmt.Printf("slice %v\n", A)
}
