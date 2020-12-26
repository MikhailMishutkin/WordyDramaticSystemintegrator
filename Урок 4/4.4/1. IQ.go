package main

import "fmt"
func main() {
	fmt.Println("IQ космонавтов")
	fmt.Println("===================")
 	
  fmt.Println("Введите IQ первого космонавта")
	var valueIq1 int
	fmt.Scan(&valueIq1)

  fmt.Println("Введите IQ второго космонавта")
	var valueIq2 int
	fmt.Scan(&valueIq2)

  fmt.Println("Введите IQ третьего космонавта")
	var valueIq3 int
	fmt.Scan(&valueIq3)

  if valueIq1 > valueIq2 && valueIq1 > valueIq3 {
    fmt.Println("Коммандиром экипажа назначается космонавт под номером 1")
  } else if valueIq2 > valueIq1 && valueIq1 > valueIq3 {
    fmt.Println("Коммандиром экипажа назначается космонавт под номером 2")
  } else if valueIq1 <= 0 || valueIq2 <= 0 || valueIq3 <= 0 {
    fmt.Println("Все введённые значения IQ должны быть больше нуля!!!")
  } else {
    fmt.Println("Коммандиром экипажа назначается космонавт под номером 3")
  }
}