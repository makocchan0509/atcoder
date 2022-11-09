package main

import (
	"fmt"
	"math/big"
)

func main() {
	var A, B, C, D, E, F int64
	fmt.Scan(&A, &B, &C, &D, &E, &F)

	biga := big.NewInt(A)
	bigb := big.NewInt(B)
	bigc := big.NewInt(C)
	bigd := big.NewInt(D)
	bige := big.NewInt(E)
	bigf := big.NewInt(F)

	ab := biga.Mul(biga, bigb)
	abc := ab.Mul(ab, bigc)

	de := bigd.Mul(bigd, bige)
	def := de.Mul(de, bigf)

	all := abc.Sub(abc, def)
	r := big.NewInt(998244353)

	ans := all.Mod(all, r)
	fmt.Println(ans)
}
