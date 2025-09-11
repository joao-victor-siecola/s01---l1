package main

import "fmt"

// Exercício 04 - Verificador de Login
func verificarLogin(usuario, senha string) bool {
    if usuario == "admin" && senha == "senha" {
        return true
    }
    return false
}

func main() {
    for {
        var usuario, senha string
        fmt.Print("Digite o usuário: ")
        fmt.Scanln(&usuario)
        fmt.Print("Digite a senha: ")
        fmt.Scanln(&senha)

        if verificarLogin(usuario, senha) {
            fmt.Println("Login bem-sucedido!")
            break
        } else {
            fmt.Println("Usuário ou senha incorretos. Tente novamente.")
        }
    }
}
