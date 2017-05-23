package services

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	i := 1

	arr := []string{"item", "item", "item", "item", "item", "item"}

	for _, name := range arr {
		switch name {
		case "item":
			if i == 3 {
				continue
			}
			i++
		}
	}

	fmt.Println(i)
	if i != 4 {
		t.Error("We expet i to be 4")
	}
}
