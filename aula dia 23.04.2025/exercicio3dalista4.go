package main

import "fmt"

func main() {
	var nums [10]int
	var pares []int
	var impares []int
	somaPares := 0

	fmt.Print("Digite 10 números:")
	for i := range nums {
		fmt.Scan(&nums[i])
		if nums[i]%2 == 0 {
			pares = append(pares, nums[i])
			somaPares += nums[i]
		} else {
			impares = append(impares, nums[i])
		}
	}

	fmt.Print("a) Números pares:", pares)
	fmt.Print("b) Soma dos números pares:", somaPares)
	fmt.Print("c) Números ímpares:", impares)
	fmt.Print("d) Quantidade de números ímpares:", len(impares))
}
