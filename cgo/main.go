package main

//#include <stdio.h>
//void callC(){
// 		printf("Printing C statement 1\n");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("Printing Go statement 1...")
	C.callC()
	fmt.Println("Printing Go statement 2 ...")
}
