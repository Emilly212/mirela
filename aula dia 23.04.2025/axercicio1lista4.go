package main

import "fmt"

func main() {
	nums := []int{10, 20, 55, 30, 65, 40, 75, 50, 85, 90}

	encontrou := false
	for i, num := range nums {
		if num > 50 {
			if !encontrou {
				fmt.Println("Números superiores a 50:")
				encontrou = true
			}
			fmt.Printf("Número: %d, Posição: %d\n", num, i)
		}
	}

	if !encontrou {
		fmt.Println("Não existe nenhum número superior a 50.")
	}
}
