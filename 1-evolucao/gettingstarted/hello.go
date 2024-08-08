package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// ! Rodar Padrão:
// go run .

// ! Baixar Modulo:
// go mod tidy .
// When you ran go mod tidy, it located and downloaded the
// rsc.io/quote module that contains the package you imported.
// By default, it downloaded the latest version -- v1.5.2.
