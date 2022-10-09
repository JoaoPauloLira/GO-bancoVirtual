package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta contas.VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	contaDoGuilherme := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome: "Guilherme",
		},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	contaDoGuilherme.Depositar(125.5)

	contaDaBruna := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Bruna"}, NumeroConta: 222, NumeroAgencia: 111222}
	contaDaBruna.Depositar(200)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	//Ponteiro
	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = clientes.Titular{Nome: "Cris"}
	contaDaCris.Depositar(500)

	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris)

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = clientes.Titular{Nome: "Silvia"}
	contaDaSilvia.Depositar(500)

	fmt.Println(contaDaSilvia.ObterSaldo())

	fmt.Println(contaDaSilvia.Sacar(400))
	fmt.Println(contaDaSilvia.ObterSaldo())

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	fmt.Println(contaDaSilvia.ObterSaldo())
	fmt.Println(contaDaSilvia.Depositar(2000))

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	contaSilvia := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Silvia"}}
	contaSilvia.Depositar(300)
	contaDoGustavo := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Gustavo"}}
	contaDoGustavo.Depositar(100)

	status := contaDoGustavo.Transferir(200, &contaSilvia)

	fmt.Println(status)
	fmt.Println(contaSilvia)
	fmt.Println(contaDoGustavo)

	status2 := contaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(status2)
	fmt.Println(contaSilvia)
	fmt.Println(contaDoGustavo)

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	//contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
	//	Nome: "Bruno",
	//	CPF: "123.111.123.12"
	//	Profissao: "Desenvolvedor"},
	//	NumeroAgencia:123, NumeroConta: 123456, ObterSaldo():100}

	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroConta: 123, NumeroAgencia: 123456}
	contaDoBruno.Depositar(100)
	fmt.Println(contaDoBruno)

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)

	fmt.Println(contaDoDenis)

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("")

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())

}
