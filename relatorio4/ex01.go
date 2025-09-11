package main

import "fmt"

// Exercício 01 - Classificador de Números
func classificar_numero(num int) string {
    if num > 0 {
        return "Positivo"
    } else if num < 0 {
        return "Negativo"
    } else {
        return "Zero"
    }
}

func main() {
    var num int
    fmt.Print("Digite um número para classificar: ")
    fmt.Scanln(&num)
    fmt.Println("Classificação:", classificar_numero(num))
}
