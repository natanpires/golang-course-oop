package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

func main() {
	var conta *ContaCorrente
	conta = new(ContaCorrente)
	conta.titular = "Natan"
	conta.numeroConta = 234
	conta.numeroAgencia = 1001
	conta.saldo = 1000.1
	fmt.Println(conta.saldo)
	fmt.Println(conta.Sacar(500))
	fmt.Println(conta.saldo)
}
