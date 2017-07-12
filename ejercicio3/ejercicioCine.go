package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Comprador 1
type Comprador interface {
	GetDto() float32
	SetDto(d float32)
}

//Cliente 2
type Cliente struct {
	dto  float32
	tipo string
}

func (c *Cliente) GetDto() float32 {
	return c.dto
}

func (c *Cliente) SetDto(new float32) {
	c.dto = new
}

func (c *Cliente) getTipo() string {
	return c.tipo
}

type Entrada interface {
	Precio(precioGeneral float64) float64
}

//EntradaGeneral  con dto adicional
type EntradaGeneral struct {
	dto  float64
	tipo string
}

//Precio Jubilado
func (e *EntradaGeneral) Precio(precioGeneral float64) float64 {
	return precioGeneral * (1 - e.dto)
}

//EntradaJubilado  con dto adicional
type EntradaJubilado struct {
	EntradaGeneral
}

//EntradaInvitado  con dto adicional
type EntradaInvitado struct {
	EntradaGeneral
}

//Funcion container de Entradas
type Funcion struct {
	entradas []Entrada
}

func (f *Funcion) addTicket(ent Entrada) {
	f.entradas = append(f.entradas, ent)
}

func (f *Funcion) getEarnings(precioGeneral float64) float64 {
	var res = 0.0
	for i := 0; i < len(f.entradas); i++ {
		res += f.entradas[i].Precio(precioGeneral)
	}
	return res
}

//Cine container de funciones
type Cine struct {
	funciones     []Funcion
	precioGeneral float64
}

func (c *Cine) addFunction(f *Funcion) {
	c.funciones = append(c.funciones, *f)
}

func (c *Cine) getFunctionEarnings(index int) float64 {
	var f = c.funciones[index]
	return f.getEarnings(c.precioGeneral)
}

func (c *Cine) addTicketToFuncion(funcion int, ent Entrada) {
	c.funciones[funcion].addTicket(ent)
}

func main() {
	fmt.Println("Ejercicio de Cine")
	fmt.Printf("Ingrese el precio general de la Entrada: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	//fmt.Println("Lei: ", input)
	precioEntrada, _ := strconv.ParseFloat(input, 2)
	cine := &Cine{make([]Funcion, 10, 10), precioEntrada}
	fmt.Println(cine.precioGeneral)
	ent := &EntradaGeneral{0, "General"}
	var ent2 = &EntradaJubilado{}
	ent2.dto = 0.5
	ent2.tipo = "Jubilado"

	var ent3 = &EntradaInvitado{}
	ent3.dto = 1
	ent3.tipo = "Invitado"
	fmt.Printf("Entrada para %s $ %v\n", ent2.tipo, ent2.Precio(cine.precioGeneral))

	fmt.Printf("Entrada para %s $ %v\n", ent.tipo, ent.Precio(cine.precioGeneral))

	fmt.Printf("Entrada para %s $ %v\n", ent3.tipo, ent3.Precio(cine.precioGeneral))

	cine.addFunction(&Funcion{})
	cine.addTicketToFuncion(0, &EntradaGeneral{0, "General"})
	cine.addTicketToFuncion(0, ent2)
	cine.addTicketToFuncion(0, ent3)
	var total = cine.getFunctionEarnings(0)
	fmt.Printf("Total de Funcion %v: %v\n", 0, total)

}
