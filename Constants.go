package main

import ("fmt"
		"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 500000000
	fmt.Println(math.Sin(n))

	const d = 3e20 / n
	fmt.Println(d)
}

