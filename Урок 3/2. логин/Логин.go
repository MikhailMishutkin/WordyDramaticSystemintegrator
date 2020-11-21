package main

import "fmt"

func main() {
	fmt.Println("Приветствуем, пожалуйста, пройдите регистрацию")
	fmt.Println("Введите логин")
  var log string
  fmt.Scan(&log)
	fmt.Println("Введите пароль")
	var pass string
	fmt.Scan(&pass)
	fmt.Println("Укажите ваш возраст")
	var age int
	fmt.Scan(&age)
  fmt.Println("===================")
	fmt.Println("Поздравляю, ", log, ",", "вы успешно зарегистрированы!" )
	fmt.Println("Ваш пароль ", pass)
	fmt.Println("Ваш возраст: ", age)
}
