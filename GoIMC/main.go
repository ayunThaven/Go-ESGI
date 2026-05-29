package main

import "fmt"

func main() {
	var poids float64
	var taille float64

	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	fmt.Print("Entrez votre poids en kg : ")
	fmt.Scanln(&poids)

	fmt.Print("Entrez votre taille en m : ")
	fmt.Scanln(&taille)

	imc := poids / (taille * taille)

	fmt.Printf("IMC : %.2f\n", imc)

	switch {
	case imc < IMCMaigreur:
		fmt.Println("Catégorie : Maigreur")
	case imc < IMCNormal:
		fmt.Println("Catégorie : Normal")
	case imc < IMCSurpoids:
		fmt.Println("Catégorie : Surpoids")
	default:
		fmt.Println("Catégorie : Obésité")
	}
}
