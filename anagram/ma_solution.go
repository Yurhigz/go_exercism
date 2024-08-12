package anagram
import (
    "strings"
    "sort"
)
func StringToRuneSlice(s string) []rune {
      var r []rune
      for _, runeValue := range s {
              r = append(r, runeValue)
      }
      return r
}

func SortStringByCharacter(s string) string {
      r := StringToRuneSlice(s)
      sort.Slice(r, func(i, j int) bool {
              return r[i] < r[j]
      })
      return string(r)
}

func Detect(subject string, candidates []string) []string {
    results := []string{}
	for _,word := range candidates {
        if  SortStringByCharacter(strings.ToLower(subject)) == SortStringByCharacter(strings.ToLower(word)) && strings.ToLower(subject) != strings.ToLower(word) {
            results = append(results,word)
        }
    }
    return results
}
