package main

import "fmt"
func main() {
	fmt.Println("Модуль числа")
	fmt.Println("===================")
 	fmt.Println("Введите любое целое число:")
	var number int
	fmt.Scan(&number)

  if number < 0 {
     number = -number
  }
  fmt.Println("Модуль числа:", number)
}