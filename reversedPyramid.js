// Pseudocode:
// FUNCTION printReversedPyramid(rows):
//     FOR i FROM 0 TO rows - 1 DO:
//         SET spaces TO ' ' REPEATED i TIMES
//         SET stars TO '*' REPEATED ((rows - i) * 2 - 1) TIMES
//         PRINT spaces + stars
// END FUNCTION

function printReversedPyramid(rows) {
    for (let i = 0; i < rows; i++) {
        let spaces = ' '.repeat(i);
        let stars = '*'.repeat((rows - i) * 2 - 1);
        console.log(spaces + stars);
    }
}

// Example usage
printReversedPyramid(5);
