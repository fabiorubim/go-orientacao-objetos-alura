package main

//c "banco/contas" //apelido para o umport seria "c"
import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// contaDaSilvia := contas.ContaCorrente{}
	// contaDaSilvia.Titular = "Silvia"
	// contaDaSilvia.Saldo = 500

	// fmt.Println(contaDaSilvia.Saldo)
	// fmt.Println(contaDaSilvia.Sacar(10))
	// fmt.Println(contaDaSilvia.Saldo)
	// status, valor := contaDaSilvia.Depositar(500)
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia.Saldo)

	// contaDaMaria := contas.ContaCorrente{Titular: "Maria", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// statusConta := contaDaMaria.Transferir(200, &contaDoGustavo)
	// fmt.Println(statusConta)

	// fmt.Println(contaDaMaria)
	// fmt.Println(contaDoGustavo)

	// fmt.Println()

	// contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Bruno", CPF: "123.111.112.12", Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// 	Saldo:         110}

	clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.456.789-65", Profissao: "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoBruno.Depositar(100)

	fmt.Println(contaDoBruno)
	fmt.Println(contaDoBruno.ObterSaldo())

	PagarBoleto(&contaDoBruno, 60)
	fmt.Println(contaDoBruno.ObterSaldo())

	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(500)
	PagarBoleto(&contaLuisa, 400)
	fmt.Println(contaLuisa.ObterSaldo())

}

// Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.

// Java possui um conceito muito parecido, também chamado de interface. A grande diferença é que, em Go, um tipo implementa uma interface implicitamente.
