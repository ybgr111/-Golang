package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Solution(A *[]int) int {

	N := len(*A)
	var result int

	sort.Ints(*A)

	for i := 0; i < N; i = i + 2 {
		if i == N-1 {
			result = (*A)[i]
			break
		}
		if (*A)[i] != (*A)[i+1] {
			result = (*A)[i]
			break
		}
	}

	return result
}

func main() {

	var Str string
	fmt.Print("Enter N:")
	fmt.Fscan(os.Stdin, &Str)

	N, err := strconv.Atoi(Str)
	if err != nil {
		fmt.Print("Invalid type")
		return
	}

	if N < 1 || N > 1000000 || N%2 == 0 {
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

		if A[i] < 1 || A[i] > 1000000000 {
			fmt.Print("Invalid number")
			return
		}
	}

	fmt.Printf("%v\n", Solution(&A))

}
