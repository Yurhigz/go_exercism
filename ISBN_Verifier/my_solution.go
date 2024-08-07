package isbn
import (
    "strings"
    "regexp"
)

func IsValidISBN(isbn string) bool {
    total := 0
    trimmed := strings.ReplaceAll(isbn,"-","")
    re := regexp.MustCompile(`(\d{9}X)|(\d{10})`)
	if re.MatchString(trimmed) && len(trimmed) == 10 {
            for i := 9 ; i>= 0 ; i-- {
				if trimmed[9-i] == 'X' {
                    total += 10 * (i+1)
                }else  {
                    total += int(trimmed[9-i]- '0') * (i+1)
                } 
            }
    	return total % 11 == 0
    }
	return false

}
