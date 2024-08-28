package go_random

import (
	"fmt"
	"testing"
)

func TestRandom_MustRandomString(t *testing.T) {
	fmt.Println(MustRandomString(16))
}

func TestRandom_MustRandomNumberStr(t *testing.T) {
	fmt.Println(MustRandomNumberStr(16))
}

func TestRandom_MustRandomBytes(t *testing.T) {
	fmt.Println(MustRandomBytes(16))
}

func TestRandom_MustRandomFloat64(t *testing.T) {
	fmt.Println(MustRandomFloat64(0.003, 0.123))
}

func TestRandomCountFromSlice(t *testing.T) {
	values := RandomCountFromSlice([]string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}, 3)
	fmt.Println(values)
}
