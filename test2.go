package main

import (
	"fmt"
)

func main() {
	/*var vie int = 12 // déclaration de la variable vie
	fmt.Println(vie) // affichage de la variable vie*/

	/*var vie, argent, puissance int // déclaration de plusieurs variables sur une seule ligne
	fmt.Println("vie : ", vie)
	fmt.Println("argent : ", argent)
	fmt.Println("puissance : ", puissance)*/

	/*var vie, argent, puissance int = 10, 20, 30 // surchargement des variables
	fmt.Println("vie : ", vie)
	fmt.Println("argent : ", argent)
	fmt.Println("puissance : ", puissance)*/

	/*var vie int = 20
	var nom string = "Default"
	var vitesse float32 = 5.4

	fmt.Println("vie : ", vie)
	fmt.Println("nom : ", nom)
	fmt.Println("vitesse : ", vitesse)*/

	/*var (
		vie     int     = 20
		nom     string  = "Default"
		vitesse float32 = 5.4
	)

	fmt.Println("vie : ", vie)
	fmt.Println("nom : ", nom)
	fmt.Println("vitesse : ", vitesse)*/

	/*
	   Déclaration des variables dynamiques
	*/
	/*flt := 15.5   //  sera automatiquement de type float
	in := 5       //  sera automatiquement de type int
	st := "hello" //  sera automatiquement de type string
	bol := true   //  sera automatiquement de type boolean

	fmt.Printf("Le type de la varialbe flt est %T\n", flt)
	fmt.Printf("Le type de la varialbe in est %T\n", in)
	fmt.Printf("Le type de la varialbe st est %T\n", st)
	fmt.Printf("Le type de la varialbe bol est %T\n", bol)*/

	/*const maConstante int = 50 // déclaration d'une constante

	fmt.Println("ma Constante : ", maConstante)*/

	const maConstante int = 50

	maConstante = 50

	fmt.Println("ma Constante : ", maConstante)

}
