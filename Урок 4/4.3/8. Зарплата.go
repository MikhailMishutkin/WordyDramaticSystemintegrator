package main

import "fmt"
func main() {
	fmt.Println("Средняя разница между зарплатами")
	fmt.Println("===================")
 	
  var min int // переменная для минимального значения зарплаты
  var middleSalary int // переменная для средней зарплаты
  var max int

  fmt.Println("Введите первую зарплату")
	var salary1 int
	fmt.Scan(&salary1)

  fmt.Println("Введите вторую зарплату")
	var salary2 int
	fmt.Scan(&salary2)

  fmt.Println("Введите третью зарплату")
	var salary3 int
	fmt.Scan(&salary3)


  if salary1 * 2 > salary2 + salary3 {
    max = salary1
    Println("Минимальное значение у числа", salary2)
 
  } else {
    min = salary1
    Println("Минимальное значение у числа", salary1)
  }

// } else if salary1 == salary2 {
 //   Println("Числа не должны быть равны!!! Введите заново)
}