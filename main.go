package main

import "fmt"

func main() {

  fmt.Println("Симулятор маршрутки")

 passengers := 0 //количество пассажиров на момент события
 
 // вводим константы текстовых запросов и сообщений об остановках
 const (
   messageOut = "Сколько пассажиров вышло на остановке?"
   messageIn = "Сколько пассажиров зашло на остановке?"
   driveBy = "---Едем-Едем---"
   messageStart = "Начало маршрута. Остановка Сквер фон Неймана. В салоне пассажиров: "
   messageStart2 = "Отправляемся с остановки Сквер фон Неймана. В салоне пассажиров: "
   messageSecondStop = "2-я остановка: проспект Томаса Куртца. В салоне пассажиров:"
   messageSecondStop2 = "Отправляемся с остановки проспект Томаса Куртца. В салоне пассажиров: "
   messageThirdStop = "3-я остановка: переулок Аллена. В салоне пассажиров: "
   messageThirdStop2 = "Отправляемся с остановки переулок Аллена. В салоне пассажиров: "
   messageFourthStop ="4-я остановка: бульвар Алгоритмов. В салоне пассажиров: "
   messageFourthStop2 = "Отправляемся с остановки бульвар Алгоритмов. В салоне пассажиров: "
   messageLastStop = "Конец маршрута: тупик Гейтса"
 )
 fmt.Println(messageStart, passengers)
  var passengersOut int 
  fmt.Println(messageIn)
  var passengersIn int
  fmt.Scan(&passengersIn)

  total := passengersIn*20
  
  passengers += passengersIn - passengersOut
  fmt.Println(messageStart2, passengers)
  fmt.Println("У кассира", total, "руб.")

  fmt.Println(driveBy)
  fmt.Println(messageSecondStop, passengers)
  fmt.Println(messageOut)
  fmt.Scan(&passengersOut)
  fmt.Println(messageIn)
  fmt.Scan(&passengersIn)
  
  total += passengersIn*20
  passengers += passengersIn - passengersOut
  fmt.Println(messageSecondStop2, passengers)
  fmt.Println("У кассира", total, "руб.")

  fmt.Println(driveBy)
  fmt.Println(messageThirdStop, passengers)
  fmt.Println(messageOut)
  fmt.Scan(&passengersOut)
  fmt.Println(messageIn)
  fmt.Scan(&passengersIn)
  
  total += passengersIn*20
  passengers += passengersIn - passengersOut
  fmt.Println(messageThirdStop2, passengers)
  fmt.Println("У кассира", total, "руб.")

  fmt.Println(driveBy)
  fmt.Println(messageFourthStop, passengers)
  fmt.Println(messageOut)
  fmt.Scan(&passengersOut)
  fmt.Println(messageIn)
  fmt.Scan(&passengersIn)
  
  total += passengersIn*20
  passengers += passengersIn - passengersOut
  fmt.Println(messageFourthStop2, passengers)
  fmt.Println("У кассира", total, "руб.")

  fmt.Println(driveBy)
  fmt.Println(messageLastStop)
  fmt.Println("Всего заработали:", total, "руб.")
 
  fmt.Println("Зарплата водителя:", total / 4, "руб.")
  fmt.Println("Расходы на топливо: ", total / 5, "руб.")
  fmt.Println("Налоги: ", total / 5, "руб.")
  fmt.Println("Расходы на ремонт машины: " , total / 5, "руб.")
  total = total - (total / 4) - (total * 3 / 5) 
  fmt.Println("Итого:", total, "руб.")

}
