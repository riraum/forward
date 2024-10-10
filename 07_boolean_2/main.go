package main

import "fmt"

/*

More boolean exercises.

Also, note that you can use parentheses to group expressions, just like in math.

E.g.:
    false && (false || true)  // false
    (false && false) || true  // true

*/

func main() {
	// fmt.Println(false && (false || true))
	// fmt.Println((false && false) || true)
	logicalTable(true, true)
	logicalTable(true, false)
	logicalTable(false, false)
	logicalTable(false, true)
	strIntCheck("", 0)
	strIntCheck("Gopher", 3)
	strIntCheck("Git", 2)
	strIntCheck("Go", 4)
	// gradeValid(-10, 101, 101, 101, 90)
}

// TODO: Write a function that prints a truth table for all the following
// Operations:
//   - AND
//   - OR
//   - NAND (NOT AND)
//   - NOR  (NOT OR)
//   - XOR  (EXCLUSIVE OR)
//
// E.g.: logicalTable(true, true) would print:
// AND : true
// OR  : true
// NAND: false
// ...
// Call this function for all 4 combinations of true and false.
func logicalTable(bool1 bool, bool2 bool) {
	// Varible used
	// var AND bool = bool1 && bool2
	// fmt.Println(AND)
	// Hardcoded print results
	// if bool1 == true && bool2 == true {
	// 	fmpt.Println("true")
	// }
	// Direct print statements
	// AND // true, true == true || true, false == false || false, false == false || false, true == false
	fmt.Println("AND", bool1 && bool2)
	// OR // true, true == true || true, false == true || false, false == false || false, true == true
	fmt.Println("OR", bool1 || bool2)
	// NAND // true, true == false || true, false == true, || false, false == true || false, true == true
	fmt.Println("NAND", bool1 != bool2 || !bool1)
	//  NOR // true, true ==  false || true, false == false || false, false == true || false, true == false
	fmt.Println("NOR", !bool1 && !bool2)
	// // XOR // true, true == false || true, false == true || false, false == false || false, true == true
	fmt.Println("XOR", bool1 != bool2)
}

// TODO: Write a function that takes a string and an integer and performs
// the following checks:
// - If the string is empty, return "empty"
// - If the string is longer than the value of the integer
//   - If the integer is even, return "longer and even"
//   - If the integer is odd, return "longer and odd"
//
// - If the string is shorter than the integer, return "shorter"
func strIntCheck(str string, num int) {
	// var stringLength = len(str)
	// var integerLength = len(num)
	if str == "" {
		fmt.Println("empty")
		// TODO find way to get int length
	} else if len(str) > num {
		if len(str)%2 == 0 {
			fmt.Println("longer and even")
		} else {
			fmt.Println("longer and odd")
		}
	} else {
		fmt.Println("shorter")
	}
}

// TODO: Write a function that takes 5 integers (grades) and returns the
// following:
// - If any of the grades is less than 0 or greater than 100, return
// "invalid grade".
// - If the average of the grades is greater than or equal to 90
//   - If no grade is below 80, return "A+"
//   - If no grade is below 70, return "A"
//   - If any grade is below 70, return "A-"
//
// - If the average of the grades is greater than or equal to 80, return
// "B"
//   - If no grade is below 80, return "B+"
//   - If no grade is below 70, return "B"
//   - If any grade is below 60, return "B-"
//
// - If the average of the grades is greater than or equal to 70, return
// "C"
//   - If no grade is below 70, return "C+"
//   - If no grade is below 60, return "C"
//   - If any grade is below 60, return "C-"
// func calcGrade(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int) {
// 	var avg = (grade1 + grade2 + grade3 + grade4 + grade5) / 5
// 	// var allGradesAnd = grade1 && grade2 && grade3 && grade4 && grade5
// 	// var allGradesOr = grade1 || grade2 || grade3 || grade4 || grade5
// 	// if grade.. > 100 || grade.. < 0 print invalid
// 	 {
// 		fmt.Println(("invalid grade"))
// 		// If average grade  >= 90
// 	} if avg >= 90 {
// 		if grade1 < 80 && grade2 < 80 && grade3 < 80 && grade4 < 80 && grade5 < 80 {
// 			fmt.Println("A+")
// 		} else if grade1 < 70 && grade2 < 70 && grade3 < 70 && grade4 < 70 && grade5 < 70 {
// 			fmt.Println("A")
// 		} else if grade1 < 70 || grade2 < 70 || grade3 < 70 || grade4 < 70 || grade5 < 70 {
// 			fmt.Println("A-")
// 		}
// 	}

// }
// func gradeValid(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int) bool {
// 	(grade1 < 0 || grade1 > 100) || (grade2 < 0 || grade2 > 100) || (grade3 < 0 || grade3 > 100) || (grade4 < 0 || grade4 > 100) || (grade5 < 0 || grade5 > 100)
// 	return bool
// }

// TODO: Write a function that takes 3 numbers, representing the sides of a
// triangle, and returns the type of the triangle:
// - If any of the sides is less than or equal to 0, return "invalid"
// - If the sum of the lengths of the two shortest sides is less than or
// equal to the length of the longest side, return "invalid"
// - If all three sides are equal, return "equilateral"
// - If two sides are equal, return "isosceles"
// - If no sides are equal, return "scalene"
// }
