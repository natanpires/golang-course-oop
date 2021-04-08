package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.Saldo

	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso!", c.Saldo
	}
	return "Saldo insuficiente!", c.Saldo
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Deposito realizado com sucesso!", c.Saldo
	}
	return "Ocorreu um problema ao depositar!", c.Saldo
}

func (c *ContaCorrente) Transferir(valor float64, conta *ContaCorrente) bool {
	if valor > 0 && c.Saldo > valor {
		c.Saldo -= valor
		conta.Depositar(valor)
		return true
	}
	return false
}
