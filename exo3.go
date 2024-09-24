package TPLISA

// Fonction pour trier les intervalles selon leurs fin (tri à bulles)
func triIntervalle(intervals [][]int) {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] > intervals[j][1] {
				// Échange des intervalles
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}

// Foct pour trouver le nbr minimal d'intervalles à retirer
func Ft_non_overlap(intervals [][]int) int {
	// Si la liste est vide, on retourne 0
	if len(intervals) == 0 {
		return 0
	}

	// Trier les intervalles par leur point de fin
	triIntervalle(intervals)

	nbr := 0
	fin := intervals[0][1]

	// Parcourir les intervalles pour compter les chevauchements
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < fin {
			nbr++
		} else {
			fin = intervals[i][1]
		}
	}

	return nbr
}

/*func main() {
	fmt.Print(TPLISA.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Print(TPLISA.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
}*/
