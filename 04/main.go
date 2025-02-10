package main

// Descrição: Crie um programa que converte temperaturas entre Celsius e Fahrenheit.
import (
	"fmt"
	"strings"	
)

func main() {

	temperature := 0.0
	choice := ""

	fmt.Print("Digite a temperatura: \n")
	fmt.Scan(&temperature)

	fmt.Print("Digite a escala... (C para Celsius ou F para Fahrenheit)\n")
	fmt.Scan(&choice)

	upperChoiceStr := strings.ToUpper(choice)

	if upperChoiceStr == "C" {
		result := (temperature * 9/5) + 32

		fmt.Printf("%.1f°C é igual a %.1f°F\n", temperature, result)
	} else if upperChoiceStr == "F" {
		result := (temperature - 32) * 5/9

		fmt.Printf("%.1f°F é igual a %.1f°C\n", temperature, result)
		
	} else {
		fmt.Print("Opção inválida\n")
	}
	
}
