package main

import (
	"fmt"
	"math"
)

/*// déclaration de la fonction affichage()
func affichage() {
	fmt.Println("#################################")
	fmt.Println("\tBonjour")
	fmt.Println("#################################")
}

func main() {
	affichage() // appel de la fonction affichage()
}*/

/*// prend en paramètre un type string et un int
func affichage(nom string, age int) {
	fmt.Println("Bonjour", nom, "vous avez", age, "ans")
}

func main() {
	affichage("Hatim", 9)
	affichage("Alex", 12)
}*/

/*// Fonction qui retourne un type int
func maxNbr(a int, b int) int {
	if a > b {
		return a // retourne le variable a de type int
	}
	return b // retourne le variable b de type int
}
func main() {
	max := maxNbr(10, 30) // stockage du retour de la fonction dans une variable
	fmt.Println(max)

	// Utilisation directe du retour de la fonction sans la stocker dans une variable
	fmt.Printf("Valeur : %d , Type : %T\n", maxNbr(50, 30), maxNbr(50, 30))
}*/

/*func main() {
	a := 5
	b := 8
	fmt.Println("Avant fonction a =", a, " b =", b)
	a, b = additionTrois(a, b) // stockage du retour de la fonction dans deux variables
	fmt.Println("Après fonction a =", a, " b =", b)
}

// retourne deux types int
func additionTrois(a int, b int) (int, int) {
	return a + 3, b + 3
}*/

/*func main() {
	fmt.Println(addition(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13))
}

// déclaration d'une fonction avec des paramètres infinis
func addition(param ...int) int {
	total := 0
	for _, value := range param { //j'ai mis un underscore "_" car je ne souhaite pas récupérer l'index de la range
		total += value
	}
	return total
}*/

// Déclaration d'une fonction qui prend en paramètres un float64 et une fonction anonyme
func rajouterDix(a float64, fAnonyme func(float64) float64) /**/ {
	operation := fAnonyme(a) // Appel à notre fonction anonyme
	result := operation + 10
	fmt.Println(result)
}

func main() {
	// stockage de notre fonction anonyme dans une variable
	racineCarree := func(x float64) float64 { return math.Sqrt(x) }
	rajouterDix(9, racineCarree)

	/*
	   il est possible aussi d'utiliser directement une fonction anonyme
	   dans une variable sans forcement la stocker dans une variable
	*/
	rajouterDix(5, func(x float64) float64 { return math.Pow(x, 2) })
}
