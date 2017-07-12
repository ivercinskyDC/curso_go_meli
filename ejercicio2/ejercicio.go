package main

import (
	"fmt"
	"os"
	"strconv"
)

type btree struct {
	value       int
	left, right *btree
}

func (t *btree) addValue(newValue int) *btree {
	if t == nil {
		t = &btree{newValue, nil, nil}
		return t
	}

	if t.value > newValue {
		t.left = t.left.addValue(newValue)
	} else {
		t.right = t.right.addValue(newValue)
	}
	return t
}

func (t *btree) flat() []int {
	if t == nil {
		return make([]int, 0, 0)
	}
	var l = append(t.left.flat(), t.value)
	var s = append(t.right.flat(), t.value)
	for _, v := range s {
		l = append(l, v)
	}
	return l

}

func prinSlice(slice []int) {
	for _, value := range slice {
		fmt.Printf("%d ", value)
	}
}
func init() {
	fmt.Println("Ejercicio 2")
}
func main() {
	inputs := make([]int, 0)
	var arbol *btree
	if len(os.Args) < 2 {
		panic("Necesitas ingresar mas parametros")
	}
	for i := 1; i < len(os.Args); i++ {
		input := os.Args[i]
		value, err := strconv.Atoi(input)
		if err == nil {
			inputs = append(inputs, value)
			arbol = arbol.addValue(value)
		}
	}
	//fmt.Println(arbol.right.value)
	//prinSlice(inputs)
	fmt.Println("Entrada :")
	prinSlice(inputs)
	fmt.Println()
	fmt.Println("Salida: ")
	prinSlice(arbol.flat())

	fmt.Println("Fin Ejercicio 2")
}
