package main
// Descrição: Crie um programa que gera a tabuada de um número escolhido pelo usuário.
import("fmt")

func main() {

	num := 0

	fmt.Printf("Digite um número da tabuada: ")
	fmt.Scan(&num)
	fmt.Printf("\nTabuada de %d\n", num)

	for i := 1; i < 11; i++ {
		result := num * i

		fmt.Printf("%d x %d = %d\n", num, i, result)
	}
	
}