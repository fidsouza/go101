package main

import (
	"fmt"
	"log"
	"os"
)

func lancarMoeda(lado string) {
	var cara, coroa int
	switch lado {
	case "cara":
		cara++
	case "coroa":
		coroa++
	default:
		fmt.Println("Caiu em pe")
	}

}

func main() {
	a, b := 10, 10

	if a > b {
		fmt.Println(" a e maior que b")
	} else if a < b {
		fmt.Println(" a e menor que b")
	} else {
		fmt.Println("a igual a b ")
	}

	file, err := os.Open("hello.txt")

	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

	lancarMoeda("empe")

}
