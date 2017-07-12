package main

import (
	"fmt"

	"os"

	"strconv"

	"github.com/ivercinskyDC/curso/utils"
)

func init() {
	fmt.Printf("Recibi estos argumentos: %v\n", os.Args)
	if len(os.Args) < 2 {
		fmt.Println("No recibi el parametro para acumular!")
		os.Exit(1)
	}
}

func main() {
	input := os.Args[1]
	fmt.Println("Lei: ", input)
	top, _ := strconv.Atoi(input)
	res := utils.AddWithCond(utils.Mint.IsDiv3Or5, top)
	fmt.Printf("La suma de los elementos multiplos de 3 o 5 desde 1 hasta %d es = %d\n", top, res)
}
