package main

/*

Go has packages.

You have encountered packages already with the "fmt" package, which
provides functions for formatting strings.

This is a Standard Library package, which you can find documentation for at
https://pkg.go.dev/std.

E.g. https://pkg.go.dev/fmt#Println

A package is a collection of Go source files in the same directory that are
compiled together. Functions, types, variables, and constants defined in one
source file are visible to all other source files within the same package.

The package 'main' is special. It defines a standalone executable program, not
a library. Within the 'main' package, the function 'main' is also special: it's
where execution of the program begins.

See internal/geometry for an example of a package.
In the geometry packages, there are 2 types and associated methods.

The function float64Eq is a helper function for comparing floating point
numbers. It is is not available outside of the geometry package.

- Functions, Types, Variables, and Constants that start with a capital letter
  are exported from the package and are available to other packages.
  They are public.

- Functions, Types, Variables, and Constants that start with a lowercase letter
  are not exported from the package and are not available to other packages.
  They are private.

To use a package that is not in the standard library, you need to import it. You
need to use the full path to the package. This is the path from the root of the
module to the package.

If you look at 'go.mod', you can see that the module is called 'github.com/riraum/forward'.
So any import path that starts with 'github.com/riraum/forward' is a path to a package in
this module.

Note how the import path for the geometry package is
'github.com/riraum/forward/internal/geometry', which is the path from the root
of the module to the geometry package.


If you want to use an external package, you need to add it to the 'go.mod' file.
You can do this by running 'go get <package>' in the terminal.
Or you can add the import statement and then run 'go mod tidy' in the terminal.

*/

import (
	"fmt"

	"github.com/riraum/forward/internal/geometry"

	"github.com/riraum/forward/internal/person"
)

func main() {
	p := geometry.Point{X: 0, Y: 0}
	q := geometry.Point{X: 0, Y: 4}
	fmt.Println(p.Distance(q))

	// Hint: this main function should stay relatively small, all the logic
	// should be in the packages.

	// TODO: create a new file in the internal/geometry package called
	// 'triangle.go'.

	// DONE

	// TODO: In this file, create a new type called 'Triangle' with
	// fields 'A', 'B', and 'C' of type 'Point'.

	// DONE

	// TODO: Add a method to the 'Triangle' type called 'Area' that returns the
	// area of the triangle.

	// DONE

	// TODO: Add a method to the 'Triangle' type called 'Perimeter' that returns
	// the perimeter of the triangle.

	// DONE

	// TODO: Add a method to the 'Triangle' type called 'String' that returns
	// a string representation of the triangle in the following format:
	// "Triangle{
	//   A: {X: 0, Y: 0},
	//   B: {X: 0, Y: 4},
	//   C: {X: 3, Y: 0},
	//   Area: 6.0,
	//   Perimeter: 12.0,
	// }"

	// WIP

	// TODO: Create a new package called 'person'

	// DONE

	// TODO: In this package, create a new type called 'Person' with fields
	// 'Name' of type 'string' and 'Age' of type 'int', 'Height' of type
	// 'float64', and 'Weight' of type 'float64'.

	// DONE

	// TODO: Add a method to the 'Person' type called 'Greet' that returns a
	// string greeting the person by name.

	// DONE
	Anna := Person{
		Name: "Anna",
		Age: 20,
		Height: 177,
		Weight: 49,
	}
	fmt.Printf(Greet(Anna))

	// TODO: Add a method to the 'Person' type called 'BMI' that returns the
	// Body Mass Index of the person.



	// TODO: Add a method to the 'Person' type called 'IsAdult' that returns a
	// boolean indicating if the person is an adult (i.e. 18 years or older).

	// TODO: In the package 'person', create a new type called 'People' with
	// a field 'People' of type '[]Person'.
	// TODO: Add a method to the 'People' type called 'Average' that returns
	// a person representing the average age, height, and weight of the people.
	// TODO: Add a method to the 'People' type called 'Oldest' that returns
	// the oldest person.
	// TODO: Add a method to the 'People' type called 'Greets' that returns
	// a slice of strings of the people greeting each other.
	// E.g. if there are 2 people, the first person should greet the second
	// person, and the second person should greet the first person.
	// If there are 3 people, the first person should greet the second and
	// third person, and so on.

	// TODO: Use the external package https://github.com/MakeNowJust/heredoc
	// to fix the rendering of the following text.
	s := `
        the code is broken
        the indentation is wrong
        please fix the code
    `
	fmt.Println(s)
}
