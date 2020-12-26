package main

import "fmt"

func main() {
	fmt.Println("Банкомат и сдача")
	fmt.Println("===================")
 	
  fmt.Println("Введите сумму покупки")
	var purchaseSum int
	fmt.Scan(&purchaseSum)

  fmt.Println("Введите номинал первой монеты")
	var firstCoin int
	fmt.Scan(&firstCoin)

  fmt.Println("Введите номинал второй монеты")
	var secondCoin int
	fmt.Scan(&secondCoin)

  fmt.Println("Введите номинал третьей монеты")
	var thirdCoin int
	fmt.Scan(&thirdCoin)

  var noChange bool
  noChange = false


  if purchaseSum == firstCoin || purchaseSum == secondCoin || purchaseSum == thirdCoin {
    noChange = true
    } 
  if purchaseSum == firstCoin + secondCoin || purchaseSum == firstCoin + thirdCoin || purchaseSum == secondCoin + thirdCoin {
    noChange = true
  }
  if purchaseSum == firstCoin + secondCoin + thirdCoin {
    noChange = true
  }   
    
  if noChange {
    fmt.Println("Можно оплатить без сдачи")
  } else {
    fmt.Println("Не возможно оплатить без сдачи")
  }
}