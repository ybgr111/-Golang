package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Solution(A *[]int, N int) int {

	sort.Ints(*A)

	for i := 1; i <= N; i++ {
		if (*A)[i-1] != i {
			return 0
		}
	}
	return 1

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

	if N < 1 || N > 100000 {
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

	fmt.Printf("%v\n", Solution(&A, N))
}
