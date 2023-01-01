package main

import (
	"fmt"
	"strconv"
)

var (
	nome string
	n1   string
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(x, y int) int {
	return x + y
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i

	return total, err

}

func main() {
	message := "Filipe"
	hello(message)
	total, err := convertAndSum(10, "10")
	fmt.Println("Total", sum(10, 5))
	fmt.Println("Total", total, err)
}
