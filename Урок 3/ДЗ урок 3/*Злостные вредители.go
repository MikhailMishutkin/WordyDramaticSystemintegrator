package main

import "fmt"


func main() {

 fmt.Println("---Бамбук против вредителей---")
  fmt.Println("Это программа помогает рассчитать высоту и минимальное количество дней роста до заданной длинны бамбука в условиях поражения вредителями")
 fmt.Println("      ")
  fmt.Println("Введите высоту саженца:")
  var startHeight int
  fmt.Scan(&startHeight)

  fmt.Println("Введите скорость роста саженца за день:")
  var growth int
  fmt.Scan(&growth)
  
  fmt.Println("Введите прожорливость гусениц в см/день:")
  var losses int
  fmt.Scan(&losses)
  
  fmt.Println("Введите дни, отпущенные судьбой для роста саженцов:")
  var day int
  fmt.Scan(&day)


  currentHeight := (growth - losses)*day + startHeight
  fmt.Println("-----Итого-----")
  fmt.Print("Высота бамбука к концу ", day, "-ого дня будет составлять: ", currentHeight, ".\n")

  fmt.Println("      ")
  fmt.Println(" А если нужно узнать минимально необходимое количество дней роста саженцов до товарной кондиции,")
  fmt.Println(" то")
  fmt.Println("Введите необходимую длину для возможности сруба:")
  var cutHeight int
  fmt.Scan(&cutHeight)
 

  cut := (cutHeight - startHeight)/(growth - losses)
  fmt.Println("Полных дней до сруба:", cut + 1)
  }