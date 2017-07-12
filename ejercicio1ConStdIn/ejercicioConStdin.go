package main

import (
	"fmt"

	"bufio"
	"os"

	"strconv"

	"strings"

	"github.com/ivercinskyDC/curso/utils"
)

func init() {
	fmt.Printf("Ingrese el tope hasta el cual acumular: ")
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	//fmt.Println("Lei: ", input)
	top, _ := strconv.Atoi(input)
	res := utils.AddWithCond(utils.Mint.IsDiv3Or5, top)
	fmt.Printf("La suma de los elementos multiplos de 3 o 5 desde 1 hasta %d es = %d\n", top, res)
}
