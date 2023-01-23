package main

import (
	"fmt"
)

/*func test() {
	for i := 1; i < 3; i++ {
		a := 20
		a *= i
		fmt.Println(a)
	}
}

func main() {
	test()
}*/

/*func test() {
	for i := 1; i < 3; i++ {
		a := 20
		a *= i
		fmt.Println(a)
	}
	fmt.Println(a) // erreur ici
}

func main() {
	test()
}*/

/*func test() {
	a := 20 // déclaration de notre variable locale
	for i := 1; i < 3; i++ {
		a *= i
		fmt.Println("dans ma boucle for :", a)
	}
	fmt.Println("en dehors de ma boucle for :", a)
}

func main() {
	test()
}*/

/*func test() {
	a := 10
	a += 20
	fmt.Println(a)
}
func main() {
	test()
	//fmt.Println(a) // l'erreur vient d'ici
}*/

/*var g int // déclaration de notre variable globale

func test() {
	g += 20
	fmt.Println("Pendant ma fonction test() : ", g)
}
func main() {
	fmt.Println("Avant l'utilisation de la fonction test() :", g)
	test()
	fmt.Println("Pendant ma fonction main() : ", g)
	g += 30
	fmt.Println("Modifie moi encore : ", g)
}*/

var g int // déclaration de notre variable formel

func test(g int) { // déclaration de notre paramètre globale
	g += 20 // prend le dessus sur notre variable globale
	fmt.Println("Pendant ma fonction test() : ", g)
}
func main() {
	fmt.Println("Avant l'utilisation de la fonction test() :", g)
	test(20)
	fmt.Println("Pendant ma fonction main() : ", g)
	g += 30
	fmt.Println("Modifie moi encore : ", g)
}
