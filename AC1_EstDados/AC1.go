package main 

import (
	"fmt"
)

func calculaMedia(valores ...float64) float64 {
	soma := 0.0

	for _, v := range valores {
		soma += v
	}
	return soma / float64(len(valores))
}	

func verificaParidade(numero int) string {
	if numero%2 == 0 {
		return "par"
}
	return "ímpar"
}

func minhaPotencia(base, expoente int) int {
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	return resultado
}

func converteCelsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func main() {
	fmt.Println("Média:", calculaMedia(10.5, 8.3, 6.9, 7.2))
	fmt.Println("Paridade:", verificaParidade(5))
	fmt.Println("Potência:", minhaPotencia(2,3))
	fmt.Println("Celsius para Fahrenheit:", converteCelsiusParaFahrenheit(25))
}