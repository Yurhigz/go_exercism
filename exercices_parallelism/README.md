🟢 Niveau 1 : Goroutines et Channels de base
Exercice 1 : Affichage concurrent

Objectif :

    Créer deux fonctions printA() et printB(), chacune affichant respectivement "A" et "B" 10 fois.
    Exécuter les deux fonctions concurremment avec des Goroutines.

🔹 Exemple attendu : (l’ordre peut varier)

A B A B A B A B A B A B A B A B A B A B

📌 Indice : Utiliser go printA() et go printB() dans main().
Exercice 2 : Canal de communication

Objectif :

    Créer une fonction sendData(chan string) qui envoie 5 noms dans un channel.
    Dans main(), lancer sendData() dans une Goroutine et lire les données depuis le channel.

🔹 Exemple attendu :

Alice
Bob
Charlie
David
Eve

📌 Indice : Utiliser un for pour lire du channel.
Exercice 3 : Canal de synchronisation

Objectif :

    Créer une fonction worker(done chan bool) qui affiche "Travail terminé" après 2 secondes.
    Attendre la fin du worker() avant de quitter main().

📌 Indice : Utiliser done <- true et <-done pour synchroniser.
🟠 Niveau 2 : Concurrence avancée
Exercice 4 : Multiplication concurrente

Objectif :

    Créer une fonction multiply(a, b int, ch chan int) qui calcule a * b et envoie le résultat dans un channel.
    Dans main(), appeler multiply() avec trois paires de nombres en utilisant des Goroutines.
    Lire et afficher les résultats du channel.

🔹 Exemple attendu :

Résultat 1 : 15
Résultat 2 : 42
Résultat 3 : 64

📌 Indice : Créer un channel ch := make(chan int, 3) pour éviter le blocage.
Exercice 5 : Worker Pool

Objectif :

    Créer une worker pool avec 3 travailleurs qui traitent une liste de nombres.
    Chaque travailleur doit afficher "Worker X traite Y".

🔹 Exemple attendu :

Worker 1 traite 2
Worker 2 traite 4
Worker 3 traite 6
Worker 1 traite 8
Worker 2 traite 10

📌 Indice :

    Utiliser make(chan int, N) pour répartir les tâches.
    Lancer 3 workers avec go worker(id, ch).
    Fermer le channel une fois toutes les tâches envoyées.

🔴 Niveau 3 : Concurrence et synchronisation avancée
Exercice 6 : Mutex et Compteur sécurisé

Objectif :

    Créer une variable counter = 0.
    Lancer 10 Goroutines qui incrémentent counter 100 fois.
    Afficher counter après exécution.

📌 Indice :

    Utiliser sync.Mutex pour protéger counter.
    Vérifier que counter == 1000 à la fin.

Exercice 7 : Somme concurrente avec WaitGroup

Objectif :

    Calculer la somme de 10 nombres en divisant le travail sur 5 Goroutines.
    Afficher la somme totale après exécution.

📌 Indice :

    Utiliser sync.WaitGroup pour attendre toutes les Goroutines.
    Accumuler la somme dans une variable partagée avec sync.Mutex.

Exercice 8 : Pipeline concurrent

Objectif :

    Créer un pipeline en 3 étapes :
        La première Goroutine envoie des nombres dans un channel.
        La seconde multiplie par 2 et envoie dans un autre channel.
        La troisième affiche les résultats.

📌 Indice :

    Utiliser deux channels : chan int entre chaque étape.

🔹 Exemple attendu :

Entrée : 1 2 3 4 5
Multiplication : 2 4 6 8 10
Sortie : 2 4 6 8 10

🟣 Niveau 4 : Cas réels en concurrence
Exercice 9 : Scraper concurrent

Objectif :

    Simuler un scraper web concurrent :
        Lancer 5 Goroutines qui simulent une requête à une URL.
        Attendre la fin de toutes les Goroutines.
        Afficher "Scraping terminé".

📌 Indice :

    Utiliser sync.WaitGroup pour attendre toutes les Goroutines.

Exercice 10 : Chat en temps réel (Simulation)

Objectif :

    Créer un système de chat avec deux utilisateurs échangeant des messages via un channel bidirectionnel.
    Chaque utilisateur envoie et reçoit des messages concurremment.

📌 Indice :

    Utiliser deux Goroutines pour gérer l’envoi et la réception.




    🛠 Exercice : Pipeline de traitement de données
Objectif :

Tu dois implémenter un pipeline en 3 étapes :

    Génération des données → Un producteur envoie des nombres dans un channel.
    Traitement des données → Un worker lit les nombres, les multiplie par 2 et les envoie dans un autre channel.
    Affichage des résultats → Un consommateur affiche les résultats.

📌 Contraintes :

    Utiliser 3 Goroutines (une pour chaque étape).
    Utiliser 2 channels pour faire passer les données.
    Assurer que toutes les Goroutines terminent correctement.
    Utiliser sync.WaitGroup pour synchroniser la fin.

    🔥 Pistes de résolution

    Créer 2 channels : un chan int pour transmettre les nombres bruts et un autre chan int pour transmettre les résultats.
    Créer 3 fonctions :
        generateNumbers(n int, ch chan int, wg *sync.WaitGroup): Génère n nombres et les envoie dans ch.
        processNumbers(chIn chan int, chOut chan int, wg *sync.WaitGroup): Multiplie chaque nombre par 2 et l'envoie dans chOut.
        printResults(ch chan int, wg *sync.WaitGroup): Lit les résultats et les affiche.
    Gérer les Goroutines avec un sync.WaitGroup pour s'assurer que tout se termine bien.
    Fermer les channels une fois que toutes les données ont été traitées.