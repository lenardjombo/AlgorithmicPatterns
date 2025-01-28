class Main {
    public static void main(String[] args) {
        printSquare(3); // Test the function with charNum = 3
    }

    public static void printSquare(int charNum) {
        if (charNum > 0) { // Check if charNum is positive
            for (int rows = 0; rows < charNum; rows++) {
                for (int cols = 0; cols < charNum; cols++) {
                    System.out.print("*");
                }
                System.out.println(); // Move to the next line after each row
            }
        } else {
            System.out.println("Invalid input. Please provide a positive integer.");
        }
    }
}

//Pseudocode
// FUNCTION main()
//     CALL printSquare(3)  // Test the function with charNum = 3
// END FUNCTION

// FUNCTION printSquare(charNum)
//     IF charNum > 0 THEN
//         FOR rows FROM 0 TO charNum - 1 DO
//             FOR cols FROM 0 TO charNum - 1 DO
//                 PRINT "*"
//             END FOR
//             PRINT NEWLINE  // Move to the next line after completing a row
//         END FOR
//     ELSE
//         PRINT "Invalid input. Please provide a positive integer."
//     END IF
// END FUNCTION
