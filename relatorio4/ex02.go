package main

import "fmt"

// Exercício 02 - Par ou Ímpar
func parOuImpar(num int) string {
    switch num % 2 {
    case 0:
        return "O número é par"
    default:
        return "O número é ímpar"
    }
}

func main() {
    var num int
    fmt.Print("Digite um número para verificar par/ímpar: ")
    fmt.Scanln(&num)
    fmt.Println(parOuImpar(num))
}
