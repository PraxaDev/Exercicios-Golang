package main
// Descrição: Desenvolva um programa que conta o número de vogais em uma palavra.
import (
	"fmt"
	"strings"
)

func main() {

	vogais := "aeiou"
	var palavra string
	contador := 0

	fmt.Printf("Digite uma palavra: ")
	fmt.Scan(&palavra)

	palavra = strings.ToLower(palavra)

	for _, letra := range palavra {
		if strings.ContainsRune(vogais, letra) {
			contador ++
		}
	}

	fmt.Printf("A palavra '%s' tem %d vogais.\n", palavra, contador)
}

