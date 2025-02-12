package main

import "sync"

// Exercice 4 : Multiplication concurrente

// Objectif :

//     Créer une fonction multiply(a, b int, ch chan int) qui calcule a * b et envoie le résultat dans un channel.
//     Dans main(), appeler multiply() avec trois paires de nombres en utilisant des Goroutines.
//     Lire et afficher les résultats du channel.

// 🔹 Exemple attendu :

// Résultat 1 : 15
// Résultat 2 : 42
// Résultat 3 : 64

// 📌 Indice : Créer un channel ch := make(chan int, 3) pour éviter le blocage.

// 1️⃣ wg.Add(1) avant chaque Goroutine → On indique qu'il y a une nouvelle Goroutine à attendre.
// 2️⃣ defer wg.Done() dans chaque Goroutine → Réduit le compteur quand elle se termine.
// 3️⃣ wg.Wait() dans main() → Attend que toutes les Goroutines aient fini.
// 4️⃣ Une fois que toutes les Goroutines ont appelé Done(), Wait() se débloque et main() continue.

func multiply(a, b int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- a * b
}

func main() {
	operations := [][2]int{
		{3, 5},
		{6, 7},
		{8, 8},
	}
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	for _, op := range operations {
		wg.Add(1)
		go multiply(op[0], op[1], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		println("Résultat :", i)
	}

}
