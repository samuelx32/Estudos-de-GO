package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Função de ativação (sigmoide)
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Estrutura de Rede Neural Simples
type NeuralNetwork struct {
	inputSize         int
	hiddenSize        int
	hiddenLayerWeights []float64
	outputLayerWeights []float64
	hiddenLayerOutput  []float64
}

// Inicializa a rede neural com pesos aleatórios
func initializeNetwork(inputSize, hiddenSize, outputSize int) *NeuralNetwork {
	rand.Seed(time.Now().UnixNano())

	network := &NeuralNetwork{
		inputSize:          inputSize,
		hiddenSize:         hiddenSize,
		hiddenLayerWeights: make([]float64, inputSize*hiddenSize),
		outputLayerWeights: make([]float64, hiddenSize*outputSize),
		hiddenLayerOutput:  make([]float64, hiddenSize),
	}

	// Inicializa pesos aleatórios
	for i := 0; i < inputSize*hiddenSize; i++ {
		network.hiddenLayerWeights[i] = rand.Float64()
	}

	for i := 0; i < hiddenSize*outputSize; i++ {
		network.outputLayerWeights[i] = rand.Float64()
	}

	return network
}

// Executa a rede neural com entrada e retorna a saída
func feedforward(network *NeuralNetwork, input []float64, output []float64) {
	// Processa a camada oculta
	for i := 0; i < network.hiddenSize; i++ {
		sum := 0.0
		for j := 0; j < network.inputSize; j++ {
			sum += input[j] * network.hiddenLayerWeights[i*network.inputSize+j]
		}
		network.hiddenLayerOutput[i] = sigmoid(sum)
	}

	// Processa a camada de saída
	for i := 0; i < len(output); i++ {
		sum := 0.0
		for j := 0; j < network.hiddenSize; j++ {
			sum += network.hiddenLayerOutput[j] * network.outputLayerWeights[i*network.hiddenSize+j]
		}
		output[i] = sigmoid(sum)
	}
}

func main() {
	// Inicializa duas redes neurais
	network1 := initializeNetwork(2, 3, 3) // Primeira rede: 2 entradas, 3 neurônios ocultos, 3 saídas
	network2 := initializeNetwork(3, 3, 2) // Segunda rede: 3 entradas (da saída da rede 1), 2 saídas

	// Entrada inicial para a Rede 1
	input1 := []float64{1.0, 0.5}
	output1 := make([]float64, 3) // Saída da Rede 1 será a entrada para a Rede 2

	// Saída final da Rede 2
	output2 := make([]float64, 2)

	// Executa feedforward na Rede 1
	feedforward(network1, input1, output1)

	// Usa a saída da Rede 1 como entrada da Rede 2
	feedforward(network2, output1, output2)

	// Exibe a saída final da Rede 2
	fmt.Printf("Output da Rede 2: %f, %f\n", output2[0], output2[1])
}
