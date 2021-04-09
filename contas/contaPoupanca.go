package contas

import "go-oop/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!", c.saldo
	}
	return "saldo insuficiente!", c.saldo
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso!", c.saldo
	}
	return "Ocorreu um problema ao depositar!", c.saldo
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
