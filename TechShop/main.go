package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	produits   []Produit
	prochainID int
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	p.ID = c.prochainID
	c.prochainID++

	c.produits = append(c.produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, errors.New("produit introuvable")
}

func (c Catalogue) TrouverParCategorie(categorie string) []Produit {
	var resultats []Produit

	for _, produit := range c.produits {
		if strings.EqualFold(produit.Categorie, categorie) {
			resultats = append(resultats, produit)
		}
	}

	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0

	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, categorie) {
			c.produits[i].Prix = c.produits[i].Prix - c.produits[i].Prix*pct/100
			nbModifies++
		}
	}

	return nbModifies
}

func (c *Catalogue) Vendre(id int, quantite int) error {
	for i := range c.produits {
		if c.produits[i].ID == id {
			if quantite <= 0 {
				return errors.New("la quantité doit être supérieure à 0")
			}

			if c.produits[i].Stock < quantite {
				return errors.New("stock insuffisant")
			}

			c.produits[i].Stock -= quantite
			return nil
		}
	}

	return errors.New("produit introuvable")
}

func (c Catalogue) Rapport() string {
	nbProduits := 0
	valeurTotale := 0.0

	for _, produit := range c.produits {
		if produit.Actif {
			nbProduits++
			valeurTotale += produit.Prix * float64(produit.Stock)
		}
	}

	return fmt.Sprintf(
		"Nombre de produits actifs : %d\nValeur totale du stock : %.2f €",
		nbProduits,
		valeurTotale,
	)
}

func afficherProduit(p Produit) {
	fmt.Printf(
		"ID: %d | %s %s | %.2f € | Stock: %d | Catégorie: %s | Actif: %t\n",
		p.ID,
		p.Marque,
		p.Nom,
		p.Prix,
		p.Stock,
		p.Categorie,
		p.Actif,
	)
}

// region Gestion espace dans les noms
func lireTexte(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	texte, _ := reader.ReadString('\n')
	return strings.TrimSpace(texte)
}

func lireInt(reader *bufio.Reader, message string) int {
	for {
		texte := lireTexte(reader, message)

		valeur, err := strconv.Atoi(texte)
		if err == nil {
			return valeur
		}

		fmt.Println("Erreur : veuillez entrer un nombre entier.")
	}
}

func lireFloat(reader *bufio.Reader, message string) float64 {
	for {
		texte := lireTexte(reader, message)

		valeur, err := strconv.ParseFloat(texte, 64)
		if err == nil {
			return valeur
		}

		fmt.Println("Erreur : veuillez entrer un nombre.")
	}
}

//endregion

func main() {
	reader := bufio.NewReader(os.Stdin)

	catalogue := Catalogue{
		prochainID: 1,
	}

	catalogue.AjouterProduit(Produit{
		Nom:       "iPhone 15",
		Marque:    "Apple",
		Prix:      969.99,
		Stock:     10,
		Categorie: "Smartphone",
		Actif:     true,
	})

	catalogue.AjouterProduit(Produit{
		Nom:       "Galaxy S24",
		Marque:    "Samsung",
		Prix:      899.99,
		Stock:     8,
		Categorie: "Smartphone",
		Actif:     true,
	})

	catalogue.AjouterProduit(Produit{
		Nom:       "MacBook Air M3",
		Marque:    "Apple",
		Prix:      1299.99,
		Stock:     5,
		Categorie: "Ordinateur portable",
		Actif:     true,
	})

	catalogue.AjouterProduit(Produit{
		Nom:       "ThinkPad X1 Carbon",
		Marque:    "Lenovo",
		Prix:      1499.99,
		Stock:     4,
		Categorie: "Ordinateur portable",
		Actif:     true,
	})

	catalogue.AjouterProduit(Produit{
		Nom:       "MX Master 3S",
		Marque:    "Logitech",
		Prix:      109.99,
		Stock:     20,
		Categorie: "Accessoire",
		Actif:     true,
	})

	for {
		fmt.Println()
		fmt.Println("===== TechShop =====")
		fmt.Println("[1] Ajouter un produit")
		fmt.Println("[2] Chercher un produit")
		fmt.Println("[3] Appliquer des soldes")
		fmt.Println("[4] Vendre un produit")
		fmt.Println("[5] Rapport")
		fmt.Println("[0] Quitter")

		choix := lireInt(reader, "Votre choix : ")

		switch choix {
		case 1:
			var p Produit

			p.Nom = lireTexte(reader, "Nom : ")
			p.Marque = lireTexte(reader, "Marque : ")
			p.Prix = lireFloat(reader, "Prix : ")
			p.Stock = lireInt(reader, "Stock : ")
			p.Categorie = lireTexte(reader, "Catégorie : ")
			p.Actif = true

			err := catalogue.AjouterProduit(p)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté avec succès.")
			}

		case 2:
			fmt.Println("[1] Chercher par ID")
			fmt.Println("[2] Chercher par catégorie")

			typeRecherche := lireInt(reader, "Votre choix : ")

			switch typeRecherche {
			case 1:
				id := lireInt(reader, "ID du produit : ")

				produit, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println("Erreur :", err)
				} else {
					afficherProduit(produit)
				}

			case 2:
				categorie := lireTexte(reader, "Catégorie : ")

				produits := catalogue.TrouverParCategorie(categorie)

				if len(produits) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie.")
				} else {
					fmt.Println("Produits trouvés :")
					for _, produit := range produits {
						afficherProduit(produit)
					}
				}

			default:
				fmt.Println("Choix invalide.")
			}

		case 3:
			categorie := lireTexte(reader, "Catégorie concernée : ")
			pct := lireFloat(reader, "Réduction en % : ")

			nb := catalogue.AppliquerReduction(categorie, pct)

			fmt.Println(nb, "produit(s) modifié(s).")

		case 4:
			id := lireInt(reader, "ID du produit : ")
			quantite := lireInt(reader, "Quantité vendue : ")

			err := catalogue.Vendre(id, quantite)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente effectuée avec succès.")
			}

		case 5:
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("Fin du programme.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
