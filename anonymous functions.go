package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(i int) {

			i *= 2
			fmt.Println("i changed inside anonymous function scope", i)
		}(i)
		fmt.Println("i outside anonymous functon remains unaffected", i)
	}

}
