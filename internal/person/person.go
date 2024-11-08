package person

import "fmt"

// TODO: In this package, create a new type called 'Person' with fields
// 'Name' of type 'string' and 'Age' of type 'int', 'Height' of type
// 'float64', and 'Weight' of type 'float64'.

type Person struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

// TODO: Add a method to the 'Person' type called 'Greet' that returns a
// string greeting the person by name.

func (p Person) Greet() string {
	return fmt.Sprintf("Hello %s", p.Name)
}

// TODO: Add a method to the 'Person' type called 'BMI' that returns the
// Body Mass Index of the person.

func (p Person) BMI() float64 {
	return p.Weight / (p.Height * p.Height) * 10000
}

// TODO: Add a method to the 'Person' type called 'IsAdult' that returns a
// boolean indicating if the person is an adult (i.e. 18 years or older).

func (p Person) IsAdult() bool {
	if p.Age >= 18 {
		return true
	} else {
		return false
	}
}

// TODO: In the package 'person', create a new type called 'People' with
// a field 'People' of type '[]Person'.

type People struct {
	People []Person
}

// TODO: Add a method to the 'People' type called 'Average' that returns
// a person representing the average age, height, and weight of the people.

// Create method
func (p People) Average() Person {
	// Get amount of person in the slice
	amount := len(p.People)
	// debug
	fmt.Println("Print length", amount)
	sumAge := 0
	sumHeight := 0
	sumWeight := 0
	// Add the Age, Height and Weight
	for _, sum := range p.People {
		sumAge += sum.Age
		// debug
		fmt.Println("Print sumAge", sumAge)
		// for _, sumActualAge := range sumAge.Age {
		// 	sumAge += Person{}
		// }
	}

	for _, sum := range p.People {
		sumHeight += int(sum.Height)
	}

	for _, sum := range p.People {
		sumWeight += int(sum.Weight)
	}
	avgAge := sumAge / amount
	avgHeight := sumHeight / amount
	avgWeight := sumWeight / amount
	// Dived by amount
	// 	avgAge := sumAge / amount
	// 	// Return Person struct? Maybe just values?
	// 	return avgAge
	avgPerson := Person{
		Name:   "AveragePerson",
		Age:    avgAge,
		Height: float64(avgHeight),
		Weight: float64(avgWeight),
	}
	return avgPerson
	// return sumAge
	// return sumHeight
	// return sumWeight
}
