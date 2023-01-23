package main

import (
	"fmt"
)

/*func main() {
	var tableauInt [10]int
	var tableauFloat [10]float32
	var tableauString [10]string
	var tableauBool [10]bool
	fmt.Println("Valeur par défaut de la variable tableauInt :", tableauInt)
	fmt.Println("Valeur par défaut de la variable tableauFloat :", tableauFloat)
	fmt.Println("Valeur par défaut de la variable tableauString :", tableauString)
	fmt.Println("Valeur par défaut de la variable tableauBool :", tableauBool)
}*/

/*func main() {
	var tableau1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("la taille de mon tableau1 :", len(tableau1))
	fmt.Println("les valeurs de mon tableau1 :", tableau1)
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	fmt.Println("jours[0] =", jours[0])
	fmt.Println("jours[1] =", jours[1])
	fmt.Println("jours[2] =", jours[2])
	fmt.Println("jours[3] =", jours[3])
	fmt.Println("jours[4] =", jours[4])
	fmt.Println("jours[5] =", jours[5])
	fmt.Println("jours[6] =", jours[6])
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	fmt.Println("jours[7] =", jours[7]) // erreur
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	fmt.Println("Dernier jour =", jours[len(jours)-1]) // taille du tableau - 1 = dernier index
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	// récupération de l'index et de la valeur
	for index, j := range jours {
		fmt.Println(j, "est le jour numéro", (index + 1), "de la semaine")
	}
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	for i := 0; i < len(jours); i++ {
		fmt.Println(jours[i], "est le jour numéro", (i + 1), "de la semaine")
	}
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[:]) // on récupère tous les éléments
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[:3]) // trois premiers éléments
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[3:])
}*/

/*func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[1:3]) //Récupérer les résultats du premier jusqu'au troisième index (exclus) :
}*/

/*func main() {
	var jours = [5]string{"Hatim", "Robert", "Inconnu", "Ahmed", "Inconnu"}

	jours[0] = "Alex" // on remplace le premier element (ici Hatim) par Alex
	fmt.Println(jours)

	jours = replaceByHatim(jours)
	fmt.Println(jours)
}

/*
   J'utilise une fonction pour vous montrer qu'il est
   possible de prendre en paramètre un tableau
   mais aussi de retourner un tableau dans une fonction
*/ /*
func replaceByHatim(jours [5]string) [5]string {
	for i, jour := range jours {
		if jour == "Inconnu" {
			jours[i] = "Hatim" // Remplacer "Inconnu" par "Hatim"
		}
	}
	return jours
}*/

/*//Tableau à deux dimensions
func main() {
	const (
		maxLigne   int = 3 // 3 sous tableaux
		maxColonne int = 4 // 4 éléments pour chaque sous tableau
	)

	var tableau [maxLigne][maxColonne]int // Création d'un tableau à double dimension

	fmt.Println(tableau)
}*/

func main() {
	const (
		maxLigne   int = 3
		maxColonne int = 3
	)

	var doubleTableau [maxLigne][maxColonne]int

	fmt.Println(doubleTableau)

	fmt.Println("----------------------")

	//modification de la ligne 3, colonne 2
	doubleTableau[2][1] = 5 // modification du 2eme élément du 3eme tableau
	fmt.Println(doubleTableau)

	chaine := "Hello"
	fmt.Println(chaine)
}
