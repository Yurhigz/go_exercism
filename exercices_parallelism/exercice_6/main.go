package main

import (
	"fmt"
	"sync"
)

// Exercice 6 : Mutex et Compteur sécurisé

// Objectif :

//     Créer une variable counter = 0.
//     Lancer 10 Goroutines qui incrémentent counter 100 fois.
//     Afficher counter après exécution.

// 📌 Indice :

//     Utiliser sync.Mutex pour protéger counter.
//     Vérifier que counter == 1000 à la fin.

type Container struct {
	mu      sync.Mutex
	counter int
}

func (c *Container) incremental(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	for i := 1; i <= 100; i++ {
		c.counter += 1
	}
	defer c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := Container{
		counter: 0,
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go c.incremental(&wg)
	}
	wg.Wait()
	fmt.Println(c.counter)
}
