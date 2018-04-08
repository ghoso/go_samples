package main

import (
	"fmt"
)

func main() {
	var intA int32 = 120
	var intB int32 = 40

	gcdN := gcd(intA, intB)

	fmt.Println("gcd(", intA, ",", intB, ") = ", gcdN)
}

func gcd(a int32, b int32) int32 {
	var d1 int32
	var d2 int32
	var mod int32

	if a < b {
		d1 = b
		d2 = a
	} else {
		d1 = a
		d2 = b
	}

	for d2 != 1 {
		mod = d1 % d2
		if mod == 0 {
			return d2
		} else {
			d1 = d2
			d2 = mod
		}
	}
	return d2
}
