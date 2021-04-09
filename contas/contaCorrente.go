package contas

import "go-oop/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.saldo

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!", c.saldo
	}
	return "saldo insuficiente!", c.saldo
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso!", c.saldo
	}
	return "Ocorreu um problema ao depositar!", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, conta *ContaCorrente) bool {
	if valor > 0 && c.saldo > valor {
		c.saldo -= valor
		conta.Depositar(valor)
		return true
	}
	return false
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
