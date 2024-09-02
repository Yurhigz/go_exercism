package isogram
import "strings"

func IsIsogram(word string) bool {
	lettersCollect := make(map[rune]int)
    for _,letter := range strings.ToLower(word) {
        if letter != '-' && letter != ' ' {
        lettersCollect[letter]++
            }
    }
    for _,v := range lettersCollect {
        if v > 1 {
            return false
        }
    }
    return true
}
