package prima

import "fmt"

func Pascal(j, i int) int {
	if j == 0 || j == i {
		return 1
	}
	return Pascal(j-1, i-1) + Pascal(j, i-1)
}

func ShowPascal(n int) {
	for i := 0; i < n; i++ {
		for x := 0; x < n-i; x++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("%2d", Pascal(j, i-1))
		}
		fmt.Println()
	}
}
