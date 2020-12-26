package main

import "fmt"
func main() {
	fmt.Println("Студенты")
	fmt.Println("===================")
 	
  fmt.Println("Введите количество студентов")
	var valueStudents int
	fmt.Scan(&valueStudents)

  fmt.Println("Введите количество групп")
	var valueGrupps int
	fmt.Scan(&valueGrupps)

  fmt.Println("Введите порядковый номер студента")
	var countNumber int
	fmt.Scan(&countNumber)

 //const (
    SalesValue1 = 10 // скидка 10% или надбавка 10%
  //  SalesValue2 = 5 // скидка 5%//
//  ) 
  
  if valueDay == 1 {
    orderPrice = orderPrice - (orderPrice / SalesValue1)
  } else {
    orderPrice = orderPrice
  }
  if valueDay == 5 && orderPrice > 10000 {
    orderPrice = orderPrice - (orderPrice / SalesValue2) 
  } else {
    orderPrice = orderPrice
  }
  if valueGuest > 5 {
    orderPrice = orderPrice + (orderPrice / SalesValue1)
  } else {
    orderPrice = orderPrice
  }
  fmt.Println("Сумма к оплате за заказ:", orderPrice)
}