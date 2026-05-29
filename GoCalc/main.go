package main

import (
	"errors"
	"fmt"
)

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a float64, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a float64, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a float64, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a float64, b float64) float64 {
			return a / b
		}
	default:
		return nil
	}
}

func operer(a float64, b float64, op string) (float64, error) {
	if op == "/" && b == 0 {
		return 0, errors.New("division par zéro impossible")
	}

	operation := creerOperation(op)

	if operation == nil {
		return 0, errors.New("opération inconnue")
	}

	resultat := operation(a, b)

	return resultat, nil
}

func main() {
	for {
		var a float64
		var b float64
		var op string

		fmt.Print("Entrez une opération : ")
		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			fmt.Println("Fin du programme")
			break
		}

		resultat, err := operer(a, b, op)

		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Résultat :", resultat)
		}
	}
}
