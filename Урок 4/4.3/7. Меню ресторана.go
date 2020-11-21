package main

import "fmt"
func main() {
	fmt.Println("Меню ресторана")
	fmt.Println("===================")
 	fmt.Println("Именно сегодня в нашем ресторане ваш ждёт сюрприз, только сегодня!!!")
  fmt.Println("А помните какой сегодня по счёту день недели?")
	var dayCount int ///номер дня недели
	fmt.Scan(&dayCount)

  if dayCount == 1 {
    fmt.Println("Ваш сюрприз это плюшка!!!") 
    } else if  dayCount == 2 {
      fmt.Println("Ваш сюрприз это пирожок!!!") 
      } else if dayCount == 3 {
       fmt.Println("Ваш сюрприз это мандаринка!!!")
         } else if dayCount == 4 {
         fmt.Println("Ваш сюрприз это фруктовый салат!!!")
            } else if dayCount == 5 {
            fmt.Println("Ваш сюрприз это кекс!!!")
              }  else if dayCount == 6 {
                fmt.Println("Ваш сюрприз это манго!!!")
                  } else if dayCount == 7 {
                  fmt.Println("Ваш сюрприз это Живой кальмар!!!")
                  } else {
                    fmt.Println("Кто не знает как правильно вводить цифру дня недели,тот остаётся без сюрприза!!!")
                  }
  fmt.Println("Кстати у нас есть ещё меню этой недели, и это борщ, котлета и компот!!!")
}