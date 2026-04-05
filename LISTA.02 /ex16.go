package main

import  f "fmt"
import  m  "math"
func main() {
    
    var a, b, c float64 
    var raizes string
    var raiz1, raiz2 float64

    f.Print("Escreva os coeficientes da função (a, b e c): ")
    f.Scan(&a, &b, &c)
    if a == 0 {
        f.Println("O coeficiente 'a' não pode ser zero!")
        return
    }
    var delta float64 = (b * b) - (4 * a * c)
    
    if delta == 0 {
        raiz1 = -b / (2 * a) 
        raizes = "RAIZ ÚNICA"
        f.Println(raizes, ":", raiz1)
        return 
        
    } else if delta > 0 {
        raiz1 = (-b + m.Sqrt(delta)) / (2 * a) 
        raiz2 = (-b - m.Sqrt(delta)) / (2 * a) 
        raizes = "RAÍZES DISTINTAS"
        
    } else {
        raizes = "RAÍZES IMAGINÁRIAS"
        f.Println(raizes)
        return
    }
    f.Println(raizes, "-> Raiz 1:", raiz1, "| Raiz 2:", raiz2)
}
