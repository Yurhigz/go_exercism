package main

import "fmt"

func sendData(c chan string) {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for _, name := range names {
		c <- name // Envoie chaque nom séparément dans le channel
	}
	close(c) // Important : On ferme le channel pour signaler qu'il n'y a plus de données
}

func main() {
	c := make(chan string)

	go sendData(c) // Démarrer la Goroutine qui envoie les données

	// Lire les données jusqu'à la fermeture du channel
	for name := range c {
		fmt.Println(name)
	}
}
