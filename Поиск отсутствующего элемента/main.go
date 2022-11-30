package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Solution(A *[]int) int {

	N := len(*A)

	sort.Ints(*A)

	if (*A)[N-1] != N+1 {
		return N + 1
	}

	for i := 1; i < N; i++ {
		if (*A)[i-1] != i {
			return i
		}
	}
	return N

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

	if N < 0 || N > 100000 {
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

		if A[i] < 0 || A[i] > 100001 {
			fmt.Print("Invalid number")
			return
		}
	}

	fmt.Printf("%v\n", Solution(&A))
}
