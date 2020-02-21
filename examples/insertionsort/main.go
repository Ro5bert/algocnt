package main

import (
	"fmt"
	"github.com/Ro5bert/algocnt"
)

func insertionSort(c *algocnt.Counter, S []int) {
	for i := 1; i < len(S); i++ {
		c.Addc(algocnt.C, "i < len(S) [success]")

		val := S[i]
		j := i - 1

		for j >= 0 && S[j] > val {
			c.Addc(algocnt.C, "j >= 0 [success]")
			c.Addc(algocnt.C, "S[j] > val [success]")
			S[j+1] = S[j]
			j--
		}
		if j >= 0 {
			c.Addc(algocnt.C, "j >= 0 [success]")
			c.Addc(algocnt.C, "S[j] > val [failure]")
		} else {
			c.Addc(algocnt.C, "j >= 0 [failure]")
		}

		S[j+1] = val
	}
	c.Addc(algocnt.C, "i < len(S) [failure]")
}

func main() {
	S := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := &algocnt.Counter{}
	fmt.Println(S)
	insertionSort(c, S)
	fmt.Println(S)
	fmt.Printf("%d comparisons.\n", c.Count(algocnt.C))

	fmt.Println()

	S = []int{5, 4, 3, 8, 1, 6, 2, 7}
	c = &algocnt.Counter{}
	fmt.Println(S)
	insertionSort(c, S)
	fmt.Println(S)
	fmt.Printf("%d comparisons.\n", c.Count(algocnt.C))
}
