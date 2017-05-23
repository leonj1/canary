package services

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	i := 0
	arr := []string{"item", "item", "item", "item", "item"}
	for idx, name := range arr {
		switch name {
		case "item":
			if idx == 3 {
				continue
			}
			i++
			fmt.Println(fmt.Sprintf("idx: %d i: %d", idx, i))
		}
	}
	if i != 4 {
		t.Errorf("We expet i to be 4, actual %d", i)
	}
}
