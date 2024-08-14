package atbash
import (
    "unicode"
    "strings"
)
// 97 - 122
func Atbash(s string) string {
    trimmed := ""
    for _,letter := range strings.ToLower(strings.ReplaceAll(s," ","")) {
        if unicode.IsLower(letter) || unicode.IsNumber(letter) {
            trimmed += string(letter)
        }
    }
    encrypted := ""
    for i := 0 ; i <= len(trimmed) - 1 ; i ++ {
        
        if i % 5 == 0 && i>0 {
            encrypted += " "
        }
        if trimmed[i] >= 97 && trimmed[i]<=122 {
        	encrypted += string(rune(219 - int(trimmed[i]))) 
        } else if trimmed[i]>=48 && trimmed[i]<=57{
            encrypted += string(trimmed[i])
        }
    }
    return encrypted
}
