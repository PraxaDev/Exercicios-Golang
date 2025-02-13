package main

// Descrição: Crie uma calculadora de Índice de Massa Corporal (IMC).
import (
	"fmt"
)

func main() {

	var peso float32
	var altura float32

	fmt.Print("Digite seu peso: ")
	fmt.Scan(&peso)

	fmt.Print("Digite sua altura: ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("\nSeu IMC é: %.1f \n", imc)

	if imc <= 18.5 {
		fmt.Print("IMC Abaixo do normal\n")
	}else if imc >= 18.6 && imc <= 24.9 {
		fmt.Print("IMC Normal\n")
	}else if imc >= 25.0 && imc <= 29.9{
		fmt.Print("IMC Sobrepeso\n")
	}else if imc >= 30.0 && imc <= 34.9{
		fmt.Print("IMC Obesidade Grau 1\n")
	}else if imc >= 35.0 && imc <= 39.9{
		fmt.Print("IMC Obesidade Grau 2\n")
	}else if imc >= 40.0 {
		fmt.Print("IMC Obesidade Grau 3\n")
	}else {
		fmt.Print("Erro")
	}

}