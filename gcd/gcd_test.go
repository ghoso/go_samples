package main

import (
	"testing"
)

func TestGcd1(t *testing.T) {
	result := gcd(100, 5)
	if result != 5 {
		t.Fatalf("failed test gcd(100,5) != %d", result)
	}
}

func TestGcd2(t *testing.T) {
	result := gcd(120, 32)
	if result != 8 {
		t.Fatalf("failed test gcd(100,5) != %d", result)
	}
}
