package main

import "fmt"
func main() {
	fmt.Println("Калькулятор скидки")
	fmt.Println("===================")
	fmt.Println("Введите стоимость первого товара: ")
	var price1 int
	fmt.Scan(&price1)
  fmt.Println("Введите стоимость второго товара: ")
	var price2 int
	fmt.Scan(&price2)
  fmt.Println("Введите стоимость третьего товара: ")
	var price3 int
	fmt.Scan(&price3)
 
  total := price1 + price2 + price3
  fmt.Println("К оплате: ", total, "руб.")
  if total > 10000 {
    total -= total/10
    fmt.Println("Вам предоставлена скидка 10%, сумма к оплате", total, "руб.")
  }
}