package main

 

import (

  "fmt"

)

 

func main() {

 a := 42

 b := 153

 

 fmt.Println("a:", a)

 fmt.Println("b:", b)

 var t1 int  // вводим дополнительные переменные t1 и t2
 var t2 int

  t1 = a
  t2 = b
  a = t2
  b = t1

  fmt.Println("Меняем местами значения переменных")
  fmt.Println("a:", a)

  fmt.Println("b:", b)
}