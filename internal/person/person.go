package person

import "fmt"

type Person struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello %s", p.Name)
}
