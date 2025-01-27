package main

import (
  "fmt"
)
func main(){
  var resultingPyramid := printPyramid(rows)
  fmt.Printf("Pyramid : %v\n",rows)
}
func printPyramid(rows int){
  //Outer loop controls number of rows
  for i := 0; i <= rows; i++{
    //First inner loop to print spaces before pyramid character
    for k := 1; k <= rows-1; k++{
      fmt.Println(" ") //Print spcaces
    }
    //Second inner loop to print pyramid characters
    for j := 1; j<= (2*i-1); j++{
       fmt.Print("*") // Print star
    }

     // Move to the next row
      fmt.Println()
  }
}

//Pseudocode
// 1. Input the number of rows, n
// 2. For each row i from 1 to n:
//    a. Print (n - i) spaces
//    b. Print (2 * i - 1) stars
//    c. Move to the next line
// 3. End
