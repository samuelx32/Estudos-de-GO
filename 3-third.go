package main

//Esse algoritmo é um resumo de algumas operações de Estatística Descritiva.

import (
    "fmt"
    "math"
)

func main() {
	var notas = []int{5, 3, 9, 3, 4, 3, 6, 3, 7, 8}

	ordenando(notas)
	fmt.Println(notas)
	
	fmt.Println("Média: ", media(notas))
	fmt.Println("Moda: ", moda(notas))
	fmt.Println("Mediana: ", mediana(notas))
    fmt.Println("Variância: ", variancia(notas, media(notas)))
    desvio_p := math.Sqrt(float64(variancia(notas, media(notas))))
    fmt.Printf("Desvio padrão: %.2f",desvio_p)
}

func variancia (notas [] int, media float32) float32{
    var soma float32
    for i := 0; i < len(notas); i++{
        if float32(notas[i]) > media{
            soma += (float32(notas[i]) - media) * (float32(notas[i]) - media)
        }else{
            soma += (media - float32(notas[i])) * (media - float32(notas[i]))
        }
        
    }
    return soma / float32(len(notas))
}

func media(notas []int) float32 {
	var soma int
	for i := 0; i < len(notas); i++ {
		soma += notas[i]
	}
	return float32(soma) / float32(len(notas))
}

func mediana(notas []int) float32 {
	meio := len(notas) / 2

	if len(notas)%2 == 0 {
		return (float32(notas[meio-1]) + float32(notas[meio])) / 2
	} else {
		return float32(notas[meio])
	}

}

func moda(notas []int) int {
	aux := 0
	aux2 := 0
	var moda int
	for i := 0; i < len(notas); i++ {
		for x := 1; x < len(notas); x++ {
			if notas[i] == notas[x] {
				aux2 += 1
			}
		}
		if aux2 > aux {
			aux = aux2
			moda = notas[i]
		}
		aux2 = 0
	}
	return moda
}

func ordenando(notas []int) int {
	var aux int
	for i := 0; i < len(notas); i++ {
		for x := 0; x < len(notas); x++ {
			if notas[i] < notas[x] {
				aux = notas[x]
				notas[x] = notas[i]
				notas[i] = aux
			}
		}
	}
	return 0
}
