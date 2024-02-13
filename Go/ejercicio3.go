package main

import (
    "fmt"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var num int

    fmt.Print("Ingrese un número para calcular su factorial: ")
    fmt.Scanln(&num)
    resultado := factorial(num)

    fmt.Println("El factorial de", num, "es igual a", resultado)
}
