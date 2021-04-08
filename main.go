package main

import (
	"fmt"
	c "go-oop/contas"
)

func main() {
	// var conta *ContaCorrente
	// conta = new(ContaCorrente)
	// conta.Titular = "Natan"
	// conta.NumeroConta = 234
	// conta.NumeroAgencia = 1001
	// conta.Saldo = 1000.1

	conta := c.ContaCorrente{
		Titular:       "Natan",
		NumeroAgencia: 1001,
		NumeroConta:   234,
		Saldo:         1000.1,
	}

	contaT := c.ContaCorrente{
		Titular:       "Joao",
		NumeroAgencia: 1001,
		NumeroConta:   502,
		Saldo:         180.,
	}

	fmt.Println(conta.Saldo, contaT.Saldo)
	fmt.Println(conta.Sacar(500))
	fmt.Println(conta.Depositar(120))
	conta.Transferir(100, &contaT)
	fmt.Println(conta.Saldo, contaT.Saldo)
}
