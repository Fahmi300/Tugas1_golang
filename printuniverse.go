package main
import "fmt"

func main() {
  var i int

  fmt.Print("Input Number: ")
  fmt.Scan(&i)
  if i == 42 {
	fmt.Println("Hello Universe")
  } else {
	fmt.Println(i)
  }
}