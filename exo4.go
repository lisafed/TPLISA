package TPLISA

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// Initialiser le prix min au premier jour et le bénéfice max à 0
	minPrix := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrix {
			minPrix = prices[i]
		} else { //on calcule le bénéfice et on met à jour le bénéfice max
			profit := prices[i] - minPrix
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

/*func main() {
	fmt.Print(TPLISA.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Print(TPLISA.Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
}*/
