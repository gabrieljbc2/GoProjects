package main

import (
    "fmt"
)

func main() {
    var num1, num2 int

    fmt.Print("Ingrese el primer número: ")
    fmt.Scanln(&num1)
    fmt.Print("Ingrese el segundo número: ")
    fmt.Scanln(&num2)

    resultado := num1 + num2

    fmt.Println("La suma de", num1, "y", num2, "es igual a", resultado)
}
