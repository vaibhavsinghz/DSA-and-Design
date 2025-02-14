package main

import (
	"fmt"
)

func main() {
	num := 1
	fmt.Printf("Bit representation of %d is %b\n", num, num)
	n2 := num << 5
	fmt.Printf("Bit representation of %d is %b\n", n2, n2)
	n3 := n2 >> 5
	fmt.Printf("Bit representation of %d is %b\n", n3, n3)

	fmt.Printf("Bit representation of %d is %b\n", 1, 0<<5) // as there is no 1 bit

	/*Implementation*/
	/*
		Since we have 32 bit in int32, so it's kind of array of size 32 which we can use as visited map
	*/
	fmt.Println("mask Implementation \n")
	mask := 0
	mask |= 1 << 5 // this will set ith bit in mask
	mask |= 1 << 2
	fmt.Printf("Bit representation of %d is %b\n", mask, mask)

	mask ^= 1 << 2 // this unset the ith bit
	// xor operator : if a and b bit is same then it gives 0 and 0 when both bit is different
	fmt.Printf("Bit representation of %d is %b\n", mask, mask)

	fmt.Println((mask >> 5) & 1) // this check if ith bit is set or not
	fmt.Println((mask >> 2) & 1)

}
