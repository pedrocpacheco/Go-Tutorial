package main

import "fmt"

var letters = []string{"A", "B", "C", "D", "E"}

func Slice() {

	var letter2 = []string{"A", "B", "C", "D"}
	fmt.Println(letter2)

	leeter3 := []string{"A", "B", "C", "D"}
	fmt.Println(leeter3)

}

// TODO [FORA DE FUNÇÃO]

// * Eu consigo declarar fora de uma função com a estrutura
// var "nomeVariavel" = []string{"A", "B", "C", "D"}

// ! Eu não consigo declarar fora de uma função com a estrutura
// "nomeVariavel" := []string{"A", "B", "C", "D"}
// ? Porque? -> Em Go, a sintaxe := faz a declaração e inicialização da variável ao mesmo tempo, inferindo o tipo automaticamente. Contudo, isso é válido apenas dentro de funções, porque o escopo global exige uma declaração explícita com var
// ? Em resumo, a sintaxe := é restrita ao escopo local (dentro de funções)

// TODO [DENTRO DE FUNÇÃO]

// * Eu consigo declarar dentro de uma função com a estrutura
// var "nomeVariavel" = []string{"A", "B", "C", "D"}

// * Eu consigo declarar dentro de uma função com a estrutura
// "nomeVariavel" := []string{"A", "B", "C", "D"}
