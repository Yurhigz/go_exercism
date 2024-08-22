package armstrong
import "strconv"

func IsNumber(n int) bool {
    if n == 0 || n < 10 {
        return true
    }

    strNum := strconv.Itoa(n)
    lenNumber := len(strNum)
    sum := 0

    for _, v := range strNum {
        digit := int(v - '0')
        power := 1
        for i := 0; i < lenNumber; i++ {
            power *= digit
        }
        sum += power
    }

    return sum == n
}
