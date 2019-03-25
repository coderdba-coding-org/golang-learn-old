package main
import "fmt"

var x, y, z int

var array1 [10] float32;

func add(a int, b int) int {

return a + b

}

func forloop(a int, b int) {

for i := a; i <= b; i++ {

fmt.Println (i);

}

}


func main() {

fmt.Println ("Hello")

fmt.Println ("sum of 6 and 4 is ", add (6,4))

fmt.Println ("x,y,z", x, y, z)

forloop (10, 13);

array1 := {101,101.1,101.2,101,3}

fmt.Println (array1)

}

