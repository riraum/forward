package person

import "testing"

// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hello %s", p.Name)
// }

func TestGreet(t *testing.T) {
	anna := Person{
		Name:   "Anna",
		Age:    17,
		Height: 160,
		Weight: 49,
	}
	test := anna.Greet()
	goodResult := "Hello Anna"
	if test != "Hello Anna" {

		t.Errorf("anna.Greet() = %v; want %v", test, goodResult)
	}
}
