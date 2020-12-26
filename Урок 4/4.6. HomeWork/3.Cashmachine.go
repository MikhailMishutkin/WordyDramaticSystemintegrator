package main

import "fmt"
func main() {
	fmt.Println("Банкомат")
	fmt.Println("===================")
 	
  fmt.Println("Введите сумму")
	var cashNeed int
	fmt.Scan(&cashNeed)

  const (
    canTake = 100 //купюра к выдаче
    canGive = 100000 // максимальная сумма выдачи
  )
  
  if cashNeed % canTake != 0 && canGive > cashNeed {
    fmt.Println("Ошибка! Введите сумму кратную 100-та")
    }  else if canGive < cashNeed && cashNeed % canTake >= 0 {
    fmt.Println("Ошибка! Введите сумму меньше 100000")
    } else {
          fmt.Println("Заберите сумму")
    }  
}