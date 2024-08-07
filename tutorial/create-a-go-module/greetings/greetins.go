package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello %s", name)
	return message
}

// ! Nome das funções
// Nome com letra maiucula -> Função Publica
// Nome com letra minuscula -> Função Privada
// * No momento de chamar a função também precisa ser Maisculo/minusculo

// ! Operador :=
// Ele declara e inicializa uma variavel na mesma linha
