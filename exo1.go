package TPLISA

func Ft_coin(coins []int, amount int) int {
	nb := 0
	sum := 0
	if amount == 0 {
		return nb
	}
	// Trier les pièces en ordre décroissant
	for i := 0; i < len(coins)-1; i++ {
		for j := i + 1; j < len(coins); j++ {
			if coins[i] < coins[j] {
				coins[i], coins[j] = coins[j], coins[i]
			}
		}
	}
	// Parcourir toutes les pièces triées
	for i := 0; i < len(coins); i++ {
		for sum+coins[i] <= amount {
			sum += coins[i]
			nb++
			if sum == amount {
				return nb
			}
		}
	}
	// Si on n'atteint amount , on retourne -1
	return -1
}

/*func main() {
	fmt.Print(TPLISA.Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Print(TPLISA.Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Print(TPLISA.Ft_coin([]int{1}, 0))        // resultat : 0
}*/
