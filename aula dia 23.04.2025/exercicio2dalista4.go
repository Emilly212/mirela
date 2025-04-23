package main

import "fmt"

func somaVet(v []int) int {
	soma := 0
	for _, num := range v {
		soma += num
	}
	return soma
}

func main() {
	var vet1 [10]int
	var vet2 [5]int

	fmt.Print("Digite 10 números para o primeiro vetor:")
	for i := range vet1 {
		fmt.Scan(&vet1[i])
	}

	fmt.Print("Digite 5 números para o segundo vetor:")
	for i := range vet2 {
		fmt.Scan(&vet2[i])
	}

	somaVet2 := somaVet(vet2[:])

	var pares, impares []int
	for _, num := range vet1 {
		if num%2 == 0 {
			pares = append(pares, num+somaVet2)
		} else {
			impares = append(impares, num+somaVet2)
		}
	}

	fmt.Print("Primeiro vetor resultante (pares):", pares)
	fmt.Print("Segundo vetor resultante (ímpares):", impares)
}
