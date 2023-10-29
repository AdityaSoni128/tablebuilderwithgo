package main

import "fmt"

func main() {
	fmt.Println("this is version 1.0.0 from table Builder")
}

func GetTable(val int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v X %v = %v \n", val, i, val*i)
	}
}
