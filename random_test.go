package go_random

import (
	"fmt"
	"testing"
)

func TestRandom_MustRandomString(t *testing.T) {
	fmt.Println(RandomInstance.MustRandomString(16))
}

func TestRandom_MustRandomNumberStr(t *testing.T) {
	fmt.Println(RandomInstance.MustRandomNumberStr(16))
}

func TestRandom_MustRandomBytes(t *testing.T) {
	fmt.Println(RandomInstance.MustRandomBytes(16))
}

func TestRandom_MustRandomFloat64(t *testing.T) {
	fmt.Println(RandomInstance.MustRandomFloat64(0.003, 0.123))
}
