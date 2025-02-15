package main

import (
	"fmt"
	"sync"
)

// Exercice 7 : Somme concurrente avec WaitGroup

// Objectif :

//     Calculer la somme de 10 nombres en divisant le travail sur 5 Goroutines.
//     Afficher la somme totale après exécution.

// 📌 Indice :

//     Utiliser sync.WaitGroup pour attendre toutes les Goroutines.
//     Accumuler la somme dans une variable partagée avec sync.Mutex.

type SumMut struct {
	mu  sync.Mutex
	sum int
}

func (s *SumMut) Sum(v []int, wq *sync.WaitGroup) {
	defer wq.Done()
	sum := 0
	for _, e := range v {
		sum += e
	}
	s.mu.Lock()
	s.sum = sum
	defer s.mu.Unlock()

}

func main() {
	listNumber := []int{15, 25, 35, 45, 55, 65, 75, 85, 95, 105}
	s := SumMut{}

	var wq sync.WaitGroup
	chunkSize := len(listNumber) / 5 // number of workers here
	for i := 0; i < 5; i++ {
		wq.Add(1)
		go s.Sum(listNumber[i*chunkSize:(1+i)*chunkSize], &wq)
	}
	wq.Wait()

	fmt.Println(s.sum)
}
