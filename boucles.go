package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	/*for compteur := 0; compteur < 10; compteur++ {
		fmt.Println(compteur+1, ") Je ne dois frapper mes camarades de classe")
	}*/

	/*scanner := bufio.NewScanner(os.Stdin)
	var age int

	for age < 18 { // on quitte la boucle s'il est majeur
		fmt.Print("Entrez votre age : ")
		scanner.Scan()
		age, _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Bienvenue en boite de nuit !")*/

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10)

	scanner := bufio.NewScanner(os.Stdin)

	max := 20

	for true { // boucle infinie
		fmt.Print("Entrez votre nombre : ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Entrez un nombre !")
			continue // on revient au début de la boucle
		}
		if nbr > max || nbr < 0 {
			fmt.Println("Votre nombre doit être compris entre 0 et ", max, " !")
			continue
		} else if nbr == randomInt {
			fmt.Println("Bien joué !")
			break // on quitte la boucle
		} else {
			fmt.Println("Dommage !")
		}
	}
}
