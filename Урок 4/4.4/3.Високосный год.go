package main

import "fmt"
func main() {
	fmt.Println("Високосный год")
	fmt.Println("===================")
 	
  fmt.Println("Введите год")
	var year int
	fmt.Scan(&year)

  if year % 100 == 0 && year % 400 > 0 {
    fmt.Println("Введенный год имеет 365 дней в году")
  }  else if year % 4 == 0 {
    fmt.Println("Введенный год - високосный! Количество дней дней в году: 366") 
    } else {
    fmt.Println("Введенный год имеет 365 дней в году")
    }
}