package main

import "fmt"

func main() {
	nomes := []string{"filipe", "joao", "thiago"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])

	}
	loops()
	loops2()

}

func loops() {
	nomes := []string{"filipe", "joao", "thiago"}
	for _, nome := range nomes {
		fmt.Println("segunda forma", nome)

	}

}

func loops2() {
	nomes := []string{"filipe", "joao", "thiago"}
	var i int
	for i < len(nomes) {
		fmt.Println("terceira forma", nomes[i])
		i++
	}

}
