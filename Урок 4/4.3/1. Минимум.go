package main

import "fmt"
func main() {
	fmt.Println("Ищем минимум")
	fmt.Println("===================")
 	
  fmt.Println("Введите первое число")
	var val1 int
	fmt.Scan(&val1)

  fmt.Println("Введите второе число")
	var val2 int
	fmt.Scan(&val2)
  
  if val1 > val2 {
    Println("Минимальное значение у числа", val2)
  } else if val1 == val2 {
    Println("Числа не должны быть равны!!! Введите заново)
  } else { 
    Println("Минимальное значение у числа", val1)
  }
}