package main

import (
	"fmt"
	"sync"
	"time"
)

// Exercice 5 : Worker Pool

// Objectif :

//     Créer une worker pool avec 3 travailleurs qui traitent une liste de nombres.
//     Chaque travailleur doit afficher "Worker X traite Y".

// 🔹 Exemple attendu :

// Worker 1 traite 2
// Worker 2 traite 4
// Worker 3 traite 6
// Worker 1 traite 8
// Worker 2 traite 10

// 📌 Indice :

//     Utiliser make(chan int, N) pour répartir les tâches.
//     Lancer 3 workers avec go worker(id, ch).
//     Fermer le channel une fois toutes les tâches envoyées.

// func worker(id int, ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range ch {
// 		println("Worker", id, "traite", i)
// 	}
// }

// var jobs chan int = make(chan int, 5)
// var wg sync.WaitGroup
// num := 3
// for i := 0; i < num; i++ {
// 	wg.Add(1)
// 	go worker(i+1, jobs, &wg)
// }
// for i := 2; i <= 10; i += 2 {
// 	jobs <- i
// }
// close(jobs)

// wg.Wait()

// Exercice 5 bis : Téléchargements concurrents

// 📌 Objectif :
// Tu dois simuler un système de téléchargement de fichiers en parallèle en utilisant un Worker Pool.

// 📌 Instructions :

//     Créer un canal jobs contenant les fichiers à télécharger (chaîne de caractères avec le nom du fichier).
//     Créer un canal results pour stocker les fichiers traités (ex: Téléchargement terminé : fichierX).
//     Démarrer un Worker Pool avec 3 workers, qui prennent chacun un fichier, simulent un téléchargement (ex: time.Sleep(1s)) et envoient un message dans results.
//     Dans main(), attendre que tous les fichiers soient téléchargés et afficher les résultats.
//     Fermer results une fois tous les jobs terminés.

// 💡 Indices :

//     Utilise make(chan string, N) pour le canal des fichiers.
//     Utilise sync.WaitGroup pour attendre la fin des workers.
//     Simule le téléchargement avec time.Sleep(1 * time.Second).
//     Ferme results après que tous les jobs ont été traités.

// 	Bonus 🎯

//     Ajoute un temps de téléchargement aléatoire (rand.Intn(3)) pour chaque fichier.
//     Ajoute un id aléatoire aux fichiers pour simuler de vrais noms.

func worker(id int, jobs chan string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %v télécharge %v \n", id, job)
		time.Sleep(1 * time.Second) // Simulation d'un téléchargement
		results <- fmt.Sprintf("Téléchargement terminé : %v", job)
	}
}

func main() {
	var wg sync.WaitGroup
	listFiles := []string{"fichier1", "fichier2", "fichier3", "fichier4", "fichier5"}
	jobs := make(chan string)
	results := make(chan string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	go func() {
		for _, file := range listFiles {
			jobs <- file
		}
		close(jobs) // Fermeture après avoir envoyé tous les fichiers
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	// Lire les résultats
	for element := range results {
		fmt.Println(element)
	}

}
