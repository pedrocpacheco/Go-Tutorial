package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	i := a[0]
	fmt.Println(i)
}

// ! Criando arrays

// * Criando array da forma normal:
// var a [4]int
// * Atribuindo valor de array
// a[0] = 1

// * Criando array e ja defindo valores
// b := [2]string{"Penn", "Teller"}

// * Criando array e computador entendendo o tamanho
// b := [...]string{"Penn", "Teller"}
