package contas

import "auxilio_bancario/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor <= c.saldo && valor > 0

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com  sucesso"
	}

	return "Saldo insuciente"

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

	return "O valor do depósito é inválido", c.saldo

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}

	return false

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
