package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := sumOf(18, 10)
	if err != nil {
		fmt.Println("there is an error", err)
	} else {
		fmt.Println(total)
	}

}

func sumOf(start, end int) (int, error) {
	if start > end {
		return 0, errors.New("start is greater than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i

	}
	return total, nil
}
