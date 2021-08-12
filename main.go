package main

import (
 "fmt"
)

func main() {
saldo := 0
operacao := "debito"
valor := 10

saldo,erro:=calculasaldo(operacao, valor, saldo)
if ( erro != "") {fmt.Println(erro)}else{fmt.Println(saldo)}

operacao = "credito"
valor = 5
saldo,erro=calculasaldo(operacao, valor, saldo)
if ( erro != "") {fmt.Println(erro)}else{fmt.Println(saldo)}

operacao = "debito"
valor = 2
saldo,erro=calculasaldo(operacao, valor, saldo)
if ( erro != "") {fmt.Println(erro)}else{fmt.Println(saldo)}

operacao = "credito"
valor = 20
saldo,erro=calculasaldo(operacao, valor, saldo)


if ( erro != "") {fmt.Println(erro)}else{fmt.Println(saldo)}

}

func calculasaldo(operacao string, valor int, saldoatual int) (int,string){
if (operacao == "debito"){
return saldoatual + valor,""}

if (operacao == "credito"){
if (valor <= saldoatual){
return saldoatual - valor,""}else {return saldoatual,"Cê tá sem dinheiro tião"}
}

return saldoatual,""
