package main

import "fmt"
func main() {
	fmt.Println("Баллы ЕГЭ")
	fmt.Println("===================")
 	
  fmt.Println("Введите оценку за первый предмет")
	var markFist int
	fmt.Scan(&markFist)

  fmt.Println("Введите оценку за второй предмет")
	var markSecond int
	fmt.Scan(&markSecond)

  fmt.Println("Введите оценку за третий предмет")
	var markThird int
	fmt.Scan(&markThird)

  total := markFist + markSecond + markThird

  if total >= 275 && total <= 300 {
    fmt.Println("Поздравляем! Вы набрали необходимое количество баллов и поступили!!!")
    }  else if total > 300 {
    fmt.Println("Будьте внимательнее! Где-то Вы ввели неверные значения") 
    } else {
    fmt.Println("Вы провалили экзамен")
    }
}