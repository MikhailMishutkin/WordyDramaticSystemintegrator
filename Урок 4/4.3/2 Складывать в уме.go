package main

import "fmt"
func main() {
	fmt.Println("Складываем в уме")
	fmt.Println("===================")
 	
  fmt.Println("Введите первое целое число")
	var val1 int
	fmt.Scan(&val1)

  fmt.Println("Введите второе целое число")
	var val2 int
	fmt.Scan(&val2)
  
  fmt.Println("Введите ваш результат сложения")
	var sum int
	fmt.Scan(&sum)


  if sum != val1 + val2 {
    sum = val1 + val2
    fmt.Println("НЕВЕРНО! На самом деле верный ответ", sum)
  } else { 
    fmt.Println("Верно! У вас хороший навык сложения в уме")
  }
}