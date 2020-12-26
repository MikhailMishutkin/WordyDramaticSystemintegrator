package main

import "fmt"
func main() {
	fmt.Println("Угадай число")
	fmt.Println("===================")
 	
  fmt.Println("Загадайте число от 1 до 5")
 	
  fmt.Println("Введите 1, если число не равно 2, или введите 0")
	var answer int
	fmt.Scan(&answer)

    if answer == 0  {
    fmt.Println("Вы загадали число 2")
	    } else {
      fmt.Println("Введите 1, если число меньше 2, или введите 0")
      fmt.Scan(&answer)
          if answer == 1 {
        fmt.Println("Вы загадали число 1")  
            } else {
         fmt.Println("Введите 1, если число не равно 4, или введите 0")
          fmt.Scan(&answer) 
              if answer == 0 {
            fmt.Println("Вы загадали число 4") 
                } else {
                 fmt.Println("Введите 1, если число меньше 4, или введите 0") 
                 fmt.Scan(&answer)
                    if answer == 1 {
                      fmt.Println("Вы загадали число 3")
                    } else {
                      fmt.Println("Вы загадали число 5")
                    }
                }
        }
      }
}