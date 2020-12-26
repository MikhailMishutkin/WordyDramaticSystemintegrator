package main

import "fmt"
func main() {
	fmt.Println("Скидка в ресторане")
	fmt.Println("===================")
 	
  fmt.Println("Введите номер дня недели")
	var valueDay int
	fmt.Scan(&valueDay)

  fmt.Println("Введите сумму заказа")
	var orderPrice int
	fmt.Scan(&orderPrice)

  fmt.Println("Введите количество гостей")
	var valueGuest int
	fmt.Scan(&valueGuest)

  const (
    SalesValue1 = 10 // скидка 10% или надбавка 10%
    SalesValue2 = 5 // скидка 5%
  ) 
  
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