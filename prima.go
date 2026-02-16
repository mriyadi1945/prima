package prima

import "fmt"

func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func ShowPrima(n int) {
	count := 0
	for i := range n {
		if isPrima(i) {
			fmt.Printf("%3d", i)
			count++
			if count == 5 {
				fmt.Println()
				count = 0
				continue
			}
		}
	}
}
