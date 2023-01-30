// Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := big.NewInt(0).SetString("1208925819614629174706176", 10)               // 2^80
	b, _ := big.NewInt(0).SetString("147808829414345923316083210206383297601", 10) // 3^80
	var tmp big.Int
	fmt.Println("a + b = ", tmp.Add(a, b))
	fmt.Println("a - b = ", tmp.Add(a, tmp.Neg(b)))
	fmt.Println("b - a = ", tmp.Add(tmp.Neg(a), b))
	fmt.Println("a * b = ", tmp.Mul(a, b))
	fmt.Println("a / b = ", tmp.Quo(a, b))
	fmt.Println("b / a = ", tmp.Quo(b, a))
}
