package main

import (
	"fmt"
	"math/big"
)

func count40elem() *big.Int {
	f := []*big.Int{big.NewInt(1), big.NewInt(3)}
	A := []*big.Int{big.NewInt(1), big.NewInt(3)}

	for len(A) < 40 {
		nextVal := new(big.Int).Mul(big.NewInt(5), f[len(f)-1])
		nextVal.Add(nextVal, f[len(f)-2])
		f = append(f, nextVal)
		if new(big.Int).Mod(nextVal, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
			A = append(A, nextVal)
		}
	}

	return A[39]
}

func main() {
	res := count40elem()
	fmt.Println(res)
}
