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

  
  if valueDay > 5 || orderPrice < 10000 {
    fmt.Println("К сожалению, скидки не предусмотрены")
  } else {
    SalesValue := orderPrice / 10 
    fmt.Println("Cумма скидки равна", SalesValue)
  } 
  
}