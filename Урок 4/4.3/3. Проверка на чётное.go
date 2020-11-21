package main

import "fmt"
func main() {
	fmt.Println("Чёт-нечет")
	fmt.Println("===================")
 	
  fmt.Println("Введите число")
	var value int
	fmt.Scan(&value)

  check := value % 2
  if check > 0 {
    fmt.Println("Введённое число", value, "- нечётное.")
  } else { 
    fmt.Println("Введённое число ", value, "- чётное.")
  }
}