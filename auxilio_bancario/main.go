package main

import (
	"auxilio_bancario/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	contaDoLicas := contas.ContaPoupanca{}
	contaDoLicas.Depositar(100)
	PagarBoleto(&contaDoLicas, 60)

	fmt.Println(contaDoLicas.ObterSaldo())

}
