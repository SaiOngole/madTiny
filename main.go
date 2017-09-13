package main

import "fmt"

func main() {
	encoded := Encode(10000)
	fmt.Println(encoded)
	decoded := Decode(encoded)
	fmt.Println(decoded)
}
