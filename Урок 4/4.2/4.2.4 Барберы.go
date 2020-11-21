package main

import "fmt"
func main() {
	fmt.Println("Барберы")
	fmt.Println("===================")
 	fmt.Println("Неухоженная борода - огромная проблема!!!")
  fmt.Println("Сколько бородачей нуждается в уходе в вашем городе?")
	var cit int
	fmt.Scan(&cit)

  fmt.Println("Сколько барберов УЖЕ сейчас спасают женские души от ран в вашем городе?")
  var barbNow int
	fmt.Scan(&barbNow)

  barbNeed := cit / 240 //Вычисляем минимальное необходимое количество барберов

  if barbNow < barbNeed {
     fmt.Println("Ситуация критическая, нехватка барберов!!!")
  }
  if barbNow >= barbNeed {
     fmt.Println("Всё ок!!! Женщины счастливы")
  }
}