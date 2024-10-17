package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (i int,err error) {
    if n <= 0 {
        err = fmt.Errorf("N must be greater than 0")
        return
    }
    for n != 1 {
			if n % 2 == 0 {
                n = n / 2
        	} else {
            	n = 3*n + 1
        		}
        	i++
        	}
    return 
}
