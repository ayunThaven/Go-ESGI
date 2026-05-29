package main

import "fmt"

type Personne struct {
	Prenom string `json:"prenom"`
	Nom    string `json:"nom"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, email : %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string `json:"rue"`
	Ville      string `json:"ville"`
	CodePostal string `json:"code_postal"`
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string  `json:"poste"`
	Service string  `json:"service"`
	Salaire float64 `json:"salaire"`
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employé : %s\nAdresse : %s\nPoste : %s\nService : %s\nSalaire : %.2f €",
		e.Presentation(),
		e.Format(),
		e.Poste,
		e.Service,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + e.Salaire*pct/100
}

type Etudiant struct {
	Personne
	Promo   string  `json:"promo"`
	Moyenne float64 `json:"moyenne"`
}

func (e Etudiant) MentionObtenue(mentions map[string]float64) string {
	switch {
	case e.Moyenne >= mentions["TB"]:
		return "TB"
	case e.Moyenne >= mentions["B"]:
		return "B"
	case e.Moyenne >= mentions["AB"]:
		return "AB"
	case e.Moyenne >= mentions["P"]:
		return "P"
	default:
		return "Aucune mention"
	}
}

func (e Etudiant) FicheEtudiant(mentions map[string]float64) string {
	return fmt.Sprintf(
		"Étudiant : %s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.Presentation(),
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(mentions),
	)
}

func main() {
	mentions := map[string]float64{
		"TB": 16,
		"B":  14,
		"AB": 12,
		"P":  10,
	}

	augmentations := map[string]float64{
		"Développement": 10,
		"Management":    5,
	}

	employes := []Employe{
		{
			Personne: Personne{"Jean", "Dupont", 35, "jean.dupont@mail.com"},
			Adresse:  Adresse{"12 rue des Lilas", "Paris", "75000"},
			Poste:    "Développeur",
			Service:  "Développement",
			Salaire:  2500,
		},
		{
			Personne: Personne{"Sophie", "Martin", 42, "sophie.martin@mail.com"},
			Adresse:  Adresse{"8 avenue Victor Hugo", "Lyon", "69000"},
			Poste:    "Manager",
			Service:  "Management",
			Salaire:  3200,
		},
	}

	etudiants := []Etudiant{
		{
			Personne: Personne{"Lucas", "Bernard", 21, "lucas.bernard@mail.com"},
			Promo:    "M1 Informatique",
			Moyenne:  15.5,
		},
		{
			Personne: Personne{"Emma", "Petit", 20, "emma.petit@mail.com"},
			Promo:    "L3 Informatique",
			Moyenne:  17.2,
		},
	}

	for i := range employes {
		service := employes[i].Service
		pct := augmentations[service]

		employes[i].AugmenterSalaire(pct)
	}

	fmt.Println("=== Employés ===")

	for _, employe := range employes {
		fmt.Println(employe.FicheEmploye())
		fmt.Println()
	}

	fmt.Println("=== Étudiants ===")

	for _, etudiant := range etudiants {
		fmt.Println(etudiant.FicheEtudiant(mentions))
		fmt.Println()
	}
}
