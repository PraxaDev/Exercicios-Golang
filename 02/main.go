package main
// Descrição: Crie um programa em Go que solicita ao usuário que insira dois números, converte essas entradas para o tipo float (para lidar com números decimais), realiza a soma desses números e exibe o resultado. Este exercício introduz conceitos básicos como entrada de dados, conversão de tipos e operações aritméticas simples.
import("fmt")

func main () {

	num1 := 0.0
	num2 := 0.0

	fmt.Print("Digite um número: \n")
	fmt.Scan(&num1)

	fmt.Print("Digite outro número: \n")
	fmt.Scan(&num2)

	result := num1 + num2

	fmt.Printf("A soma de %.1f + %.1f = %.1f\n", num1, num2, result)
}