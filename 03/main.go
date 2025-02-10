package main
// Descrição: Crie uma calculadora que realiza operações básicas (adição, subtração, multiplicação e divisão).
import("fmt")

func main() {

	num1 := 0.0
	num2 := 0.0

	fmt.Print("Digite um número: \n")
	fmt.Scan(&num1)

	fmt.Print("Digite outro número: \n")
	fmt.Scan(&num2)

	choice := ""

	fmt.Print("Digite o simbolo da operação... + adição, - subtração, * multiplicação e / divisão\n")
	fmt.Scan(&choice)

	if choice == "+" {
		result := num1 + num2
		fmt.Printf("O resultado de %.1f + %.1f = %.1f\n", num1, num2, result)
	} else if choice == "-" {
		result := num1 - num2
		fmt.Printf("O resultado de %.1f - %.1f = %.1f\n", num1, num2, result)
	} else if choice == "*" {
		result := num1 * num2
		fmt.Printf("O resultado de %.1f * %.1f = %.1f\n", num1, num2, result)
	} else if choice == "/" {
		result := num1 / num2
		fmt.Printf("O resultado de %.1f / %.1f = %.1f\n", num1, num2, result)
	} else {
		fmt.Print("Opção inválida\n")
	}
}