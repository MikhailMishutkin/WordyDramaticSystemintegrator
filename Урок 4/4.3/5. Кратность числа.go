package main

import "fmt"
func main() {
	fmt.Println("Крат ты мне или не крат")
	fmt.Println("===================")
 	
  fmt.Println("Введите делимое")
	var value1 int
	fmt.Scan(&value1)

  fmt.Println("Введите делитель")
	var value2 int
	fmt.Scan(&value2)

  check := value1 % value2
  if check > 0 {
    fmt.Println("Введённое число", value1, "не кратно", value2)
  } else { 
    fmt.Println("Введённое число", value1, "кратно", value2)
  }
}