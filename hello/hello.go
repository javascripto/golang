package main

import (
	"fmt"
	"reflect"
)

func main() {
	hello()
	variables1()
	variables2()
	drawMenu()
}

func hello() {
	fmt.Println("Hello, world!")
}

func variables1() {
	var name string = "John" // explicitly typed
	var lastname = "Doe" // inferred type
	fullname := name + " " + lastname // variable declaration var keyword + strings concatenation
	fullname = name + " " + lastname // re-assigned without declaration
	fmt.Println(name, lastname)
	fmt.Println(fullname)
}

func variables2() {
	var n1 float32 = 0.1
	var n2 float32 = 0.2
	n3 := 0.3
	result := n1 + n2
	fmt.Println(result)
	fmt.Println("O tipo da variavel n3 é", reflect.TypeOf(n3)) // float64
}

func drawMenu() {
	version := 1.0
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var option int
	fmt.Scanf("%d", &option) // ponteiro para variavel
	// fmt.Scan(&option) // Forma curta sem especificar formato

	fmt.Println("A opção escolhida foi a", option)
}

func input(message string) string {
	var input string
	fmt.Println(message)
	fmt.Scanf("%s", &input)
	return input
}
