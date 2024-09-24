package TPLISA

// Fonction pour trouver le nombre manquant
func Ft_missing(nums []int) int {
	n := len(nums)
	// Calcul de la somme n premiers entier !
	sum := n * (n + 1) / 2
	sum_actuelle := 0
	for _, num := range nums {
		sum_actuelle += num
	}
	return sum - sum_actuelle
}

//merci a ma licence de maths qui me sert enfin a quelque chose
//cette somme  on peut l'appliquer uniquement parce que dans l'exercice c'est bien précise qu'il,manque un UNIQUE nombre

/*func main() {
	fmt.Println(TPLISA.Ft_missing([]int{3, 1, 2}))                 // resultat : 0
	fmt.Print(TPLISA.Ft_missing([]int{0, 1}))                      // resultat : 2
	fmt.Print(TPLISA.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8
}*/
