package Tablebuilderwithgo

import "fmt"

func GetTable(val int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v X %v = %v \n", val, i, val*i)
	}
}
