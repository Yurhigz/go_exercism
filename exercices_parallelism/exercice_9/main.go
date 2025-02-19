package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Exercice 9 : Scraper concurrent

// Objectif :

//     Simuler un scraper web concurrent :
//         Lancer 5 Goroutines qui simulent une requête à une URL.
//         Attendre la fin de toutes les Goroutines.
//         Afficher "Scraping terminé".

// 📌 Indice :

//     Utiliser sync.WaitGroup pour attendre toutes les Goroutines.

func fetchUrl(url string) {

}

func main() {
	url := "https://en.wikipedia.org/wiki/Rust_(programming_language)"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
