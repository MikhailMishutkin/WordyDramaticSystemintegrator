package main

import "fmt"
func main() {
	fmt.Println("Дорога в Рязань")
	fmt.Println("===================")
  fmt.Println("Дорога в Рязань то ещё испытание, узнай сможешь ли доехать за 2 часа!")
	fmt.Println("Введите среднюю скорость движения:")
	var speed int
	fmt.Scan(&speed)

  distance := speed * 2
  fmt.Println("Вы проехали", distance, "км.")
  if distance >= 200 {
        fmt.Println("Вы приехали в Рязань!")
  }
}