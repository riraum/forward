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
	fmt.Println(calcGrade(-10, 101, 101, 101, 90))
	fmt.Println(calcGrade(90, 90, 90, 90, 90))
	fmt.Println(calcGrade(90, 90, 90, 90, 79))
	fmt.Println(calcGrade(90, 90, 90, 90, 69))
	fmt.Println(calcGrade(80, 80, 80, 80, 80))
	fmt.Println(calcGrade(80, 80, 80, 80, 79))
	fmt.Println(calcGrade(80, 80, 80, 80, 59))
	fmt.Println(calcGrade(70, 70, 70, 70, 70))
	fmt.Println(calcGrade(70, 70, 70, 70, 69))
	fmt.Println(calcGrade(70, 70, 70, 70, 59))
	fmt.Println(checkTriangleType(0, 0, 0))
	fmt.Println(checkTriangleType(5, 6, 11))
	fmt.Println(checkTriangleType(5, 5, 5))
	fmt.Println(checkTriangleType(5, 5, 6))
	fmt.Println(checkTriangleType(5, 4, 3))
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
	fmt.Println("NAND", !bool1 || !bool2)
	// Alternative
	fmt.Println("NAND", !(bool1 && bool2))
	//  NOR // true, true ==  false || true, false == false || false, false == true || false, true == false
	fmt.Println("NOR", !(bool1 || bool2))
	// // XOR // true, true == false || true, false == true || false, false == false || false, true == true
	fmt.Println("XOR", (bool1 || bool2) && (!bool1 || !bool2))
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
//
//   - If no grade is below 70, return "C+"
//
//   - If no grade is below 60, return "C"
//
//   - If any grade is below 60, return "C-"

// if grade.. > 100 || grade.. < 0 return invalid
func gradeValid(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int) bool {
	if (grade1 < 0 || grade1 > 100) || (grade2 < 0 || grade2 > 100) || (grade3 < 0 || grade3 > 100) || (grade4 < 0 || grade4 > 100) || (grade5 < 0 || grade5 > 100) {
		return false
	} else {
		return true
	}
}

// check if any grade is < x
func anyGradeUnder(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int, limit int) bool {
	if grade1 < limit || grade2 < limit || grade3 < limit || grade4 < limit || grade5 < limit {
		return true
	} else {
		return false
	}
}

// check if all grades are >= x
func allGradesOver(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int, limit int) bool {
	if grade1 >= limit && grade2 >= limit && grade3 >= limit && grade4 >= limit && grade5 >= limit {
		return true
	} else {
		return false
	}
}

// check if grade avg >= limit
func gradeAvgOver(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int, limit int) bool {
	if (grade1+grade2+grade3+grade4+grade5)/5 >= limit {
		return true
	} else {
		return false
	}
}

func calcGrade(grade1 int, grade2 int, grade3 int, grade4 int, grade5 int) string {
	if gradeValid(grade1, grade2, grade3, grade4, grade5) == false {
		return "invalid grade"
	}
	if gradeAvgOver(grade1, grade2, grade3, grade4, grade5, 90) == true {
		if allGradesOver(grade1, grade2, grade3, grade4, grade5, 80) == true {
			return "A+"
		} else if allGradesOver(grade1, grade2, grade3, grade4, grade5, 70) == true {
			return "A"
		} else if anyGradeUnder(grade1, grade2, grade3, grade4, grade5, 70) == true {
			return "A-"
		}
	}
	if gradeAvgOver(grade1, grade2, grade3, grade4, grade5, 80) == true {
		if allGradesOver(grade1, grade2, grade3, grade4, grade5, 80) == true {
			return "B+"
		} else if allGradesOver(grade1, grade2, grade3, grade4, grade5, 70) == true {
			return "B"
		} else if anyGradeUnder(grade1, grade2, grade3, grade4, grade5, 60) == true {
			return "B-"
		}
	}
	if gradeAvgOver(grade1, grade2, grade3, grade4, grade5, 70) == true {
		if allGradesOver(grade1, grade2, grade3, grade4, grade5, 70) == true {
			return "C+"
		} else if allGradesOver(grade1, grade2, grade3, grade4, grade5, 60) == true {
			return "C"
		} else if anyGradeUnder(grade1, grade2, grade3, grade4, grade5, 60) == true {
			return "C-"
		}
	}
	return "error"
}

// TODO: Write a function that takes 3 numbers, representing the sides of a
// triangle, and returns the type of the triangle:
// - If any of the sides is less than or equal to 0, return "invalid"
// - If the sum of the lengths of the two shortest sides is less than or
// equal to the length of the longest side, return "invalid"
// - If all three sides are equal, return "equilateral"
// - If two sides are equal, return "isosceles"
// - If no sides are equal, return "scalene"

// any side <= 0 == true
func anySideLess(side1 int, side2 int, side3 int) bool {
	if side1 <= 0 || side2 <= 0 || side3 <= 0 {
		return true
	} else {
		return false
	}
}

// sum of length of 2 shortest <= longest side == true
func sumLength(side1 int, side2 int, side3 int) bool {
	if side1 <= side2 && side3 > side1 && side3 > side2 || side2 <= side1 && side3 > side1 && side3 > side2 {
		if side1+side2 <= side3 {
			return true
		}
	} else if side3 <= side1 && side2 > side3 && side2 > side1 || side1 <= side3 && side2 > side3 && side2 > side1 {
		if side3+side1 <= side2 {
			return true
		}
	} else if side2 <= side3 && side1 > side2 && side1 > side3 || side3 <= side2 && side1 > side2 && side1 > side3 {
		if side2+side3 <= side1 {
			return true
		}
	}
	return false
}

// all sides equal == true
func allSidesEqual(side1 int, side2 int, side3 int) bool {
	if side1 == side2 && side3 == side1 {
		return true
	} else {
		return false
	}
}

// 2 sides equal == true
func twoSidesEqual(side1 int, side2 int, side3 int) bool {
	if side1 == side2 || side2 == side3 || side1 == side3 {
		return true
	} else {
		return false
	}
}

// no sides equal == true
func noSidesEqual(side1 int, side2 int, side3 int) bool {
	if side1 != side2 || side2 != side3 || side1 != side3 {
		return true
	} else {
		return false
	}
}

// master calc function
func checkTriangleType(side1 int, side2 int, side3 int) string {
	if anySideLess(side1, side2, side3) == true || sumLength(side1, side2, side3) == true {
		return "invalid"
	} else if allSidesEqual(side1, side2, side3) == true {
		return "equilateral"
	} else if twoSidesEqual(side1, side2, side3) == true {
		return "isosceles"
	} else if noSidesEqual(side1, side2, side3) == true {
		return "scalene"
	} else {
		return "error"
	}
}
