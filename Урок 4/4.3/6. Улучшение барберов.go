package main

import "fmt"
func main() {
	fmt.Println("Барберы")
	fmt.Println("===================")
 	fmt.Println("Неухоженная борода - огромная проблема!!!")
  fmt.Println("Сколько бородачей нуждается в уходе в вашем городе?")
	var cit int ///количество бородачей
	fmt.Scan(&cit)

  fmt.Println("Сколько барберов УЖЕ сейчас спасают женщин от эстетических травм в вашем городе?")
  var barbNow int
	fmt.Scan(&barbNow)
  
  citLast := cit % 240 // есть ли  неохваченный остаток бородачей
  barbNeed := cit / 240 //Вычисляем минимальное необходимое количество барберов
  if citLast > 0 {
    barbNeed++
  }
  fmt.Println(barbNeed, "это сколько героев необходимо")
  
   if barbNow < barbNeed {
     fmt.Println("Ситуация критическая, нехватка барберов!!!")
    } else {
     fmt.Println("Всё ок!!! Женщины счастливы")
  }
}