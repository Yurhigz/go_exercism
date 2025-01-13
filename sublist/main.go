package main

import "fmt"

type Relation string

const (
    SUBLIST   Relation = "sublist"
    SUPERLIST Relation = "superlist"
    EQUAL     Relation = "equal"
    UNEQUAL   Relation = "unequal"
)

func Sublist(l1, l2 []int) Relation {
    len1, len2 := len(l1), len(l2)
    
    // Vérifie si les listes sont égales
    if len1 == len2 && isSublist(l1, l2) {
        return EQUAL
    }
    
    // Vérifie si l1 est une sous-liste de l2
    if len1 < len2 && isSublist(l1, l2) {
        return SUBLIST
    }
    
    // Vérifie si l2 est une sous-liste de l1
    if len1 > len2 && isSublist(l2, l1) {
        return SUPERLIST
    }
    
    return UNEQUAL
}

// Fonction helper pour vérifier si needle est une sous-liste de haystack
func isSublist(needle, haystack []int) bool {
    if len(needle) == 0 {
        return true
    }
    if len(haystack) == 0 {
        return false
    }
    
    for i := 0; i <= len(haystack)-len(needle); i++ {
        matches := true
        for j := range needle {
            if needle[j] != haystack[i+j] {
                matches = false
                break
            }
        }
        if matches {
            return true
        }
    }
    return false
}
func main() {
	fmt.Println(isSublist([1,2,3], [1,2,3,4,5]))

}