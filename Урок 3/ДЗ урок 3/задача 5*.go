package main

import "fmt"


func main() {

 a := 42

 b := 153

  
 

 fmt.Println("a:", a)

 fmt.Println("b:", b)

  a = b * a // для значений отличных от 0
  b = a / b
  a = a / b


  fmt.Println("Меняем местами значения переменных")
  fmt.Println("a:", a)

  fmt.Println("b:", b)
}
   