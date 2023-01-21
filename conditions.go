package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	/*scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("Entrez quelque chose : ")
	scanner.Scan()                      // lancement du scanner
	entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
	fmt.Println(entreeUtilisateur)*/

	/*scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez un nombre entier : ")
	scanner.Scan()

	nbr, _ := strconv.Atoi(scanner.Text()) // conversion du type string en int

	fmt.Printf("res : %d\n", (nbr + 6))*/

	/*scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
		fmt.Println("On essaie de m'arnaquer ? aller Dehors ! Et la prochaine entrez un entier !")
		os.Exit(2) // on quitte le programmation
	}

	if age < 17 { // vérifier si l'utilisateur à au moins 18 ans
		fmt.Println("Sortez !")
	} else { // si ce n'est pas le cas alors on l'accepte pas
		fmt.Println("Entrez :)")
	}*/

	/*scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("On essaie de m'arnaquer ? allé Dehors ! Et la prochaine entrez un entier !")
		os.Exit(2)
	}

	fmt.Print("Entrez votre prenom : ")
	scanner.Scan()
	prenom := scanner.Text()

	/*
	   On a besoin de changer la graine (générateur de nombres pseudo-aléatoires) à
	   chaque exécution de programmation sinon on obtiendra le même nombre aléatoire
	*/ /*
		rand.Seed(time.Now().UnixNano())
		radomInt := rand.Intn(2) // retourne aléatoirement soit 1 soit 0
		radomInt2 := rand.Intn(2)

		if age < 18 {
			fmt.Println("Sortez !")
		} else if prenom == "Hatim" || prenom == "hatim" {
			fmt.Println("Ah c'est toi Hatim, dehors !")
		} else if age == 18 && radomInt == 1 { // si le client est chanceux et qu'il a 18 ans
			fmt.Println("Hum vous avez de la chance je suis de bonne humeur, entrez !")
		} else if radomInt2 == 0 {
			fmt.Println("Vous n'avez vraiment pas de chance sortez !")
		} else {
			fmt.Println("Entrez :)")
		}*/

	/*scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrez un entier !")
		os.Exit(2)
	}

	switch choix { // la variable qu'on souhaite vérifier
	case 0, 1: // 1 ou 0
		println("George Boole !")
	case 7:
		println("William Van de Walle!")
	case 23:
		println("Pour certains, le nombre 23 est source de nombreux mystères, qu'en penses-tu Jim Carrey?")
	case 42:
		println("la réponse à la question ultime du sens de la vie!")
	case 666:
		println("Quand le diable veut une âme, le mal devient séduisant.")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}*/

	/*switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Il est samedi")
	case time.Sunday:
		fmt.Println("Il est dimanche")
	default:
		fmt.Println("au boulot ! Le week-end est fini")
	}*/

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrez un entier !")
		os.Exit(2)
	}

	switch {
	case choix >= 2000:
		println("Ah un 2000 !")
	case choix >= 1939 && choix <= 1945:
		println("Triste année")
	case time.Now().Weekday() == time.Sunday:
		println("Dimanche !")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}
}
