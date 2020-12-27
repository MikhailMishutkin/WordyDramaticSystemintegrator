package main

import ("fmt"
 "math")

func main() {
	fmt.Println("Решаем квадратное уравнение")
	fmt.Println("===================")
 	
  fmt.Println("Введите первый коэффициент 'a'")
	var coeffA float64
	fmt.Scan(&coeffA)

  fmt.Println("Введите второй коэффициент 'b'")
	var coeffB float64
	fmt.Scan(&coeffB)

  fmt.Println("Введите третий коэффициент 'c'")
	var coeffC float64
	fmt.Scan(&coeffC)

  var rootX1 float64 // корень уравнения
  var rootX2 float64
  quadB := math.Pow(coeffB, 2) //вычисляем квадрат коэффициента b
  

  D := quadB - 4 * coeffA * coeffC // вычисляем дискриминанту
  
  rootX1 = (-coeffB + math.Sqrt(D))/2 * coeffA
  rootX2 = (-coeffB - math.Sqrt(D))/2 * coeffA

    if D > 0  {
      fmt.Println(rootX1, rootX2)
      }
    if D == 0 {
      fmt.Println(-coeffB/2*coeffA)
    } else {
      fmt.Println("Корней нет")
    }
}