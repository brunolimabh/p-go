package main

import (
	"fmt"
	// Renomear um import
	m "math"
)

func main() {
	// Palavras para definir uma variável: CONST X VAR
	const PI float64 = 3.1415
	var raio = 3.2

	// Forma reduzida de criar uma var
	// Se for iniciado com "=", significa que ela já existe e estou atribuindo um novo valor
	// Se for iniciado com ":=", significa que ela está sendo criada
	area := PI * m.Pow(raio, 2)

	// Ao declarar uma variavel, é necessaria usa-la
	fmt.Println("A área da circunferência é: ", area)

	// É possivel criar varias variaveis de uma fez
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// True ira para E e False para F
	var e, f bool = true, false
	fmt.Println(e, f)

	// g = 2, h = False, i = "epa!"
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
