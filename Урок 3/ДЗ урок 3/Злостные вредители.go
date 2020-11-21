package main

import "fmt"


func main() {

 fmt.Println("Бамбук против вредителей")

 

 
 startHeight := 100

 growth := 50

 losses := 20

 

 currentHeight := startHeight

 currentHeight += growth

 currentHeight -= losses

 currentHeight += growth

 currentHeight -= losses

 currentHeight += growth / 2


 fmt.Println("Высота бамбука к середине треьего дня:", currentHeight)
 
 
  //currentHeight := (growth - losses)*2 + growth/2 + startHeight // одной формулой Высота в середине треьего дня

  cut := (300 - startHeight)/(growth - losses)

  fmt.Println("Полных дней до сруба:", cut + 1)// полных дней необходимо  6 (280) и полдня дорости до 300 см, однако условие поставлено так что склоняюсь к выводу, что необходимо узнать количество полных дней, необходимое для преодоления отметки 300 см. С учётом того, что переменная вмещает только целый остаток от деления, то к результату необходимо прибавить 1

}
   