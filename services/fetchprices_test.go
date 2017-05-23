package services

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	i := 1

	arr := []string{"moe", "jose", "jose", "jose", "jose", "jose", "jose"}

	for _, name := range arr {
		//fmt.Println(name)
		switch name {
		case "jose":
			if i == 3 {
				continue
			}
			fmt.Println(i)
			i++
		}
	}
}
