package main

import (
	"fmt"
	"github.com/Ro5bert/algocnt"
)

func divide(c *algocnt.Counter, S []int) ([]int, []int) {
	c.EnterScope("divide")
	defer c.ExitScope()

	h := len(S) / 2
	S1 := make([]int, h)
	S2 := make([]int, h+len(S)%2)

	for i := 0; i < len(S1); i++ {
		c.Addc(algocnt.C, "i < len(S1) [success]")
		S1[i] = S[i]
	}
	c.Addc(algocnt.C, "i < len(S1) [failure]")

	for i := 0; i < len(S2); i++ {
		c.Addc(algocnt.C, "i < len(S2) [success]")
		S2[i] = S[i+h]
	}
	c.Addc(algocnt.C, "i < len(S2) [failure]")

	return S1, S2
}

func merge(c *algocnt.Counter, S1, S2, S []int) {
	c.EnterScope("merge")
	defer c.ExitScope()

	i := 0
	j := 0

	for i < len(S1) && j < len(S2) {
		c.Addc(algocnt.C, "i < len(S1) [success]")
		c.Addc(algocnt.C, "j < len(S2) [success]")
		if S1[i] < S2[j] {
			S[i+j] = S1[i]
			i++
		} else {
			S[i+j] = S2[j]
			j++
		}
		c.Addc(algocnt.C, "S1[i] < S2[j]")
	}
	if i < len(S1) {
		c.Addc(algocnt.C, "i < len(S1) [success]")
		c.Addc(algocnt.C, "j < len(S2) [failure]")
	} else {
		c.Addc(algocnt.C, "i < len(S1) [failure]")
	}

	for i < len(S1) {
		c.Addc(algocnt.C, "i < len(S1) [success]")
		S[i+j] = S1[i]
		i++
	}
	c.Addc(algocnt.C, "i < len(S1) [failure]")

	for j < len(S2) {
		c.Addc(algocnt.C, "j < len(S2) [success]")
		S[i+j] = S2[j]
		j++
	}
	c.Addc(algocnt.C, "j < len(S2) [failure]")
}

func mergeSort(c *algocnt.Counter, S []int) {
	c.EnterScope(fmt.Sprintf("mergeSort(%v)", S))
	defer c.ExitScope()

	c.Addc(algocnt.C, "len(S) < 2")
	if len(S) < 2 {
		return
	}

	S1, S2 := divide(c, S)
	mergeSort(c, S1)
	mergeSort(c, S2)
	merge(c, S1, S2, S)
}

func main() {
	S := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := &algocnt.Counter{}
	fmt.Println(S)
	mergeSort(c, S)
	fmt.Println(S)
	fmt.Printf("%d comparisons.\n", c.Count(algocnt.C))

	fmt.Println()

	S = []int{5, 4, 3, 8, 1, 6, 2, 7}
	c = &algocnt.Counter{}
	fmt.Println(S)
	mergeSort(c, S)
	fmt.Println(S)
	fmt.Printf("%d comparisons.\n", c.Count(algocnt.C))
}
