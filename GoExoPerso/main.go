package main

import "fmt"

const (
	AbonnementInconnu = iota
	AbonnementBasic
	AbonnementStandard
	AbonnementPremium
)

func main() {
	var choix int
	var avantages []string

	fmt.Println("Choisissez votre abonnement :")
	fmt.Println("1 - Basic")
	fmt.Println("2 - Standard")
	fmt.Println("3 - Premium")
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case AbonnementPremium:
		avantages = append(avantages, "Support prioritaire")
		fallthrough
	case AbonnementStandard:
		avantages = append(avantages, "Statistiques avancées")
		fallthrough
	case AbonnementBasic:
		avantages = append(avantages, "Accès à l'application")
	default:
		fmt.Println("Abonnement inconnu")
		return
	}

	fmt.Println("Avantages inclus :")

	for i := len(avantages) - 1; i >= 0; i-- {
		fmt.Println("-", avantages[i])
	}
}
