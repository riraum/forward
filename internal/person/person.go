package person

import (
	"fmt"
)

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
	return p.Age >= 18
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
	amount := float64(len(p.People))
	sumAge := 0
	sumHeight := 0.0
	sumWeight := 0.0
	// Add the Age, Height and Weight
	for _, value := range p.People {
		sumAge += value.Age
		sumHeight += value.Height
		sumWeight += value.Weight
	}

	avgAge := sumAge / int(amount)
	avgHeight := sumHeight / amount
	avgWeight := sumWeight / amount

	avgPerson := Person{
		Name:   "AveragePerson",
		Age:    avgAge,
		Height: avgHeight,
		Weight: avgWeight,
	}
	return avgPerson
}

// TODO: Add a method to the 'People' type called 'Oldest' that returns
// the oldest person.

func (p People) Oldest() Person {
	oldestAge := 0
	oldestPerson := Person{}
	for _, value := range p.People {

		if value.Age > oldestAge {
			oldestAge = value.Age
			oldestPerson = value
		}
	}
	return oldestPerson
}

// TODO: Add a method to the 'People' type called 'Greets' that returns
// a slice of strings of the people greeting each other.
// E.g. if there are 2 people, the first person should greet the second
// person, and the second person should greet the first person.
// If there are 3 people, the first person should greet the second and
// third person, and so on.

func (p People) Greet() []string {
	// Get length of Person slice in People struct
	// loop through People struct to get all names of Person
	// Greet each other Person through string with name variable (with another loop?)
	namesSlice := []string{}
	result := []string{}
	// for _, value := range names {
	for _, value := range p.People {
		namesSlice = append(namesSlice, value.Name)
	}

	for _, value := range namesSlice {
		for _, secondValue := range namesSlice {
			if value == secondValue {
				continue
			}
			result = append(result, fmt.Sprintf("Hi %v from %v!\n", value, secondValue))
		}
	}
	return result
}
