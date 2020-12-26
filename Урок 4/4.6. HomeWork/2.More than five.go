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

  
  if numberFist > 5 || numberSecond > 5 || numberThird > 5 {
    fmt.Println("Среди введённх чисел ЕСТЬ числа, большие чем 5")
    }   else {
    fmt.Println("Все введённые числа меньшие, чем 5")
    }
}