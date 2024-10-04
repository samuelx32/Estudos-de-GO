package main

import (
	"fmt"
)

// Função degrau
func funcaoDegrau(soma, limiar int) int {
	if soma >= limiar {
		return 1
	}
	return 0
}

// Função que simula o neurônio MCP
func neuronioMCP(entradas, pesos []int, limiar int) int {
	soma := 0
	for i := 0; i < len(entradas); i++ {
		soma += entradas[i] * pesos[i]
	}
	return funcaoDegrau(soma, limiar)
}

func main() {
	entradas := []int{1, 0, 1} // Entradas
	pesos := []int{2, -1, 3}   // Pesos

	limiar := 3 // Limiar

	// Calcula a saída do neurônio MCP
	saida := neuronioMCP(entradas, pesos, limiar)

	fmt.Printf("A saída do neurônio é: %d\n", saida)
}
