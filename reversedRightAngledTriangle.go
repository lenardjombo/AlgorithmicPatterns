package main

import (
	"fmt"
)

func main(){
	var num int = 5
	printReversedRightAngledTriangle(num)
}

func printReversedRightAngledTriangle(num int){
	for row := 0;row <= num;row++{
		for col := num; col >= row;col--{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//PSEUDOCODE
BEGIN
    // Define the main function
    FUNCTION main()
        // Declare and initialize the variable 'num' with the value 5
        num = 5
        
        // Call the function to print the reversed right-angled triangle
        CALL printReversedRightAngledTriangle(num)
    END FUNCTION

    // Define the function to print the reversed right-angled triangle
    FUNCTION printReversedRightAngledTriangle(num)
        // Loop through each row from 0 to 'num'
        FOR row = 0 TO num
            // Loop through each column from 'num' down to the current row
            FOR col = num DOWNTO row
                // Print an pattern
                PRINT "*"
            END FOR
            
            // Move to the next line after printing the stars for the current row
            PRINT newline
        END FOR
    END FUNCTION
END
