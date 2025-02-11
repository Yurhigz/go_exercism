package main

import (
	"fmt"
	"time"
)

// Exercice 3 : Canal de synchronisation

// Objectif :

//     Créer une fonction worker(done chan bool) qui affiche "Travail terminé" après 2 secondes.
//     Attendre la fin du worker() avant de quitter main().

// 📌 Indice : Utiliser done <- true et <-done pour synchroniser.
// 🟠 Niveau 2 : Concurrence avancée

func worker(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Travail terminé")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
