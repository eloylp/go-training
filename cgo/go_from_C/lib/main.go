package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage(m string) {
	fmt.Println(m)
}

func main() {

}
