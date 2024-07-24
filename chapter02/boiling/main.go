// Boiling exibe o ponto de ebuilição da água

package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling = %g°F or %g°C\n", f, c)
	// Saida:
	// boiling = 212°F or 100°C
}
