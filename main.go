package main

import (
	"fmt"
	"go-oop/clientes"
	c "go-oop/contas"
)

func main() {
	// var conta *ContaCorrente
	// conta = new(ContaCorrente)
	// conta.Titular = "Natan"
	// conta.NumeroConta = 234
	// conta.NumeroAgencia = 1001
	// conta.Saldo = 1000.1

	clienteNatan := clientes.Titular{
		Nome:      "Natan",
		CPF:       "000.123.123-12",
		Profissao: "Dev 2",
	}

	conta := c.ContaCorrente{
		Titular:       clienteNatan,
		NumeroAgencia: 1001,
		NumeroConta:   234,
	}

	clienteJoao := clientes.Titular{
		Nome:      "Joao",
		CPF:       "123.100.102.23",
		Profissao: "Dev 1",
	}

	contaT := c.ContaCorrente{
		Titular:       clienteJoao,
		NumeroAgencia: 1001,
		NumeroConta:   502,
	}

	conta.Depositar(1000)
	contaT.Depositar(900)
	conta.Sacar(500)
	conta.Depositar(120)
	conta.Transferir(100, &contaT)

	fmt.Println(conta.Saldo(), contaT.Saldo())
}
