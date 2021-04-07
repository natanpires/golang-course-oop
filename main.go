package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta := ContaCorrente{
		titular:       "Natan",
		numeroAgencia: 1001,
		numeroConta:   233,
		saldo:         200.31,
	}
	fmt.Println(conta)
}
