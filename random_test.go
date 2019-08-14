package go_random

import (
	"fmt"
	"testing"
)

func TestRandomClass_GetUniqueIdString(t *testing.T) {
	fmt.Println(Random.GetUniqueIdString())
}

func TestRandomClass_Test(t *testing.T) {
	fmt.Println(Random.Test())
}
