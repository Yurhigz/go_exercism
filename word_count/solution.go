package wordcount

import (
    "strings"
    "regexp"
)


type Frequency map[string]int

func WordCount(phrase string) Frequency {
    frequency := make(map[string]int)
    re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)?`)
    for _,word := range re.FindAllString(strings.ToLower(phrase),-1) {
        frequency[word] += 1
    }
    return frequency
}
