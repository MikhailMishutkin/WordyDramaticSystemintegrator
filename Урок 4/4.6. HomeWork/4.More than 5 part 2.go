package main

import "fmt"
func main() {
	fmt.Println("БОльше чем пять")
	fmt.Println("===================")
 	
  fmt.Println("Введите первое число")
	var numberFist int
	fmt.Scan(&numberFist)

  fmt.Println("Введите второе число")
	var numberSecond int
	fmt.Scan(&numberSecond)

  fmt.Println("Введите третье число")
	var numberThird int
	fmt.Scan(&numberThird)


    var quantity int         
    
    if numberFist >= 5 {
     quantity++
    } else {
      quantity = 0
    } 
    if numberSecond >= 5 {
      quantity++
    } else {
      quantity = quantity
    } 
    if numberThird >= 5 {
      quantity++
    } else {
      quantity = quantity
    }

     fmt.Println("Количество чисел больших или равных 5:", quantity)
    
}


