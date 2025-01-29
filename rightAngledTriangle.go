package main

import (
	"fmt"
)
func main(){
	var rows int = 5
	printRightAngledTriangle(rows)
}

func printRightAngledTriangle(rows int){
	//outer loop
	for i := 1;i <= rows-1;i++{
		for j := 1;j <= i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Pseudocode
// BEGIN
//     DEFINE FUNCTION printRightAngledTriangle(rows)
//         FOR i FROM 1 TO rows - 1 DO
//             FOR j FROM 1 TO i DO
//                 PRINT "*"
//             END FOR
//             PRINT NEW LINE
//         END FOR
//     END FUNCTION

//     BEGIN MAIN
//         DECLARE rows AS INTEGER = 5
//         CALL printRightAngledTriangle(rows)
//     END MAIN
// END
