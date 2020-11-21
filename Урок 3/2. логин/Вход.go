package main

import "fmt"

func main() {
	fmt.Println("Приветствуем, на нашем сайте! Пожалуйста, авторизируйтесь")
  fmt.Println("Введите логин)
  var log string
  fmt.Scan(&log)
	fmt.Println("Введите пароль")
	var pass string
	fmt.Scan(&pass)
	fmt.Println("===================")
	fmt.Println("log, ",", "вы успешно зашли!" )
}

