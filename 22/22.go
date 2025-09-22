package main

import (
	"fmt"
	"math/big"
)

/*
Большие числа и операции
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами,
но обратите внимание на возможное переполнение для ещё больших значений.
Для очень больших чисел можно использовать math/big.
*/
func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.Exp(big.NewInt(2), big.NewInt(60), nil)
	b.Exp(big.NewInt(2), big.NewInt(50), nil)

	sum := new(big.Int)
	diff := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)

	sum.Add(a, b)
	diff.Sub(a, b)
	mul.Mul(a, b)
	div.Div(a, b)
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Println()
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", diff.String())
	fmt.Printf("a * b = %s\n", mul.String())
	fmt.Printf("a / b = %s\n", div.String())
}
