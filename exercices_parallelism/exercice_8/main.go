// Exercice 8 : Pipeline concurrent

// Objectif :

//     Créer un pipeline en 3 étapes :
//         La première Goroutine envoie des nombres dans un channel.
//         La seconde multiplie par 2 et envoie dans un autre channel.
//         La troisième affiche les résultats.

// 📌 Indice :

//     Utiliser deux channels : chan int entre chaque étape.

// 🔹 Exemple attendu :

// Entrée : 1 2 3 4 5
// Multiplication : 2 4 6 8 10
// Sortie : 2 4 6 8 10

package main

import (
	"fmt"
	"sync"
)

func Sender(n int, numberChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	numberChan <- n
}

func Multiplier(nbChan chan int, multiChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range nbChan {
		multiChan <- e * 2
	}
}

func Display(multiChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range multiChan {
		fmt.Println(e)
	}
}

func main() {
	var wgSender sync.WaitGroup
	var wgMultiplier sync.WaitGroup
	var wgDisplay sync.WaitGroup

	numberChan := make(chan int)
	multiChan := make(chan int)

	wgMultiplier.Add(2)
	go Multiplier(numberChan, multiChan, &wgMultiplier)
	go Multiplier(numberChan, multiChan, &wgMultiplier)

	for i := 1; i <= 5; i++ {
		wgSender.Add(1)
		go Sender(i, numberChan, &wgSender)
	}

	// ✅ Fermer numberChan après que tous les Senders ont fini
	go func() {
		wgSender.Wait()
		close(numberChan)
	}()

	// ✅ Fermer multiChan après que tous les Multiplier ont fini
	go func() {
		wgMultiplier.Wait()
		close(multiChan)
	}()

	// 🔹 Lancer Display
	wgDisplay.Add(1)
	go Display(multiChan, &wgDisplay)

	// ✅ Attendre la fin de Display
	wgDisplay.Wait()

	wgSender.Wait()
	close(numberChan)

	wgMultiplier.Wait()
	close(multiChan)

	wgDisplay.Add(1)
	go Display(multiChan, &wgDisplay)

	wgDisplay.Wait()

}
