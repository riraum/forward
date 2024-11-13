package person

import (
	"slices"
	"testing"
)

// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hello %s", p.Name)
// }

func TestPersonGreet(t *testing.T) {
	anna := Person{
		Name:   "Anna",
		Age:    17,
		Height: 160,
		Weight: 49,
	}
	test := anna.Greet()
	wantedResult := "Hello Anna"
	if test != wantedResult {

		t.Errorf("anna.Greet() = %v; want %v", test, wantedResult)
	}
}

func TestBMI(t *testing.T) {
	anna := Person{
		Name:   "Anna",
		Age:    17,
		Height: 160,
		Weight: 49,
	}
	luna := Person{
		Name:   "Luna",
		Age:    30,
		Height: 156,
		Weight: 48,
	}
	jan := Person{
		Name:   "Jan",
		Age:    29,
		Height: 188,
		Weight: 90,
	}
	test := anna.BMI()
	wantedResult := 19.140625
	if test != wantedResult {

		t.Errorf("anna.BMI() = %v; want %v", test, wantedResult)
	}
	test2 := luna.BMI()
	wantedResult2 := 19.723865877712033
	if test2 != wantedResult2 {

		t.Errorf("luna.BMI() = %v; want %v", test2, wantedResult2)
	}
	test3 := jan.BMI()
	wantedResult3 := 25.464010864644635
	if test3 != wantedResult3 {

		t.Errorf("jan.BMI() = %v; want %v", test3, wantedResult3)
	}

}

func TestIsAdult(t *testing.T) {
	anna := Person{
		Name:   "Anna",
		Age:    17,
		Height: 160,
		Weight: 49,
	}
	luna := Person{
		Name:   "Luna",
		Age:    30,
		Height: 156,
		Weight: 48,
	}
	jan := Person{
		Name:   "Jan",
		Age:    29,
		Height: 188,
		Weight: 90,
	}
	tay := Person{
		Name:   "Tay",
		Age:    18,
		Height: 190,
		Weight: 87,
	}
	test := anna.IsAdult()
	wantedResult := false
	if test != wantedResult {

		t.Errorf("anna.IsAdult() = %v; want %v", test, wantedResult)
	}
	test2 := luna.IsAdult()
	wantedResult2 := true
	if test2 != wantedResult2 {

		t.Errorf("luna.IsAdult() = %v; want %v", test2, wantedResult2)
	}
	test3 := jan.IsAdult()
	wantedResult3 := true
	if test3 != wantedResult3 {

		t.Errorf("jan.IsAdult() = %v; want %v", test3, wantedResult3)
	}
	test4 := tay.IsAdult()
	wantedResult4 := true
	if test4 != wantedResult4 {

		t.Errorf("tay.IsAdult() = %v; want %v", test4, wantedResult4)
	}
}

func TestAverage(t *testing.T) {
	team := People{
		{Name: "Anna",
			Age:    17,
			Height: 160,
			Weight: 49,
		},
		{
			Name:   "Luna",
			Age:    30,
			Height: 156,
			Weight: 48,
		},
		{
			Name:   "Jan",
			Age:    29,
			Height: 188,
			Weight: 90,
		},
		{
			Name:   "Tay",
			Age:    33,
			Height: 190,
			Weight: 87,
		},
	}
	test := team.Average()
	wantedResult := Person{
		Name:   "AveragePerson",
		Age:    27,
		Height: 173.5,
		Weight: 68.5,
	}
	if test != wantedResult {

		t.Errorf("team.Average() = %v; want %v", test, wantedResult)
	}
}

func TestOldest(t *testing.T) {
	team := People{
		{Name: "Anna",
			Age:    17,
			Height: 160,
			Weight: 49,
		},
		{
			Name:   "Luna",
			Age:    30,
			Height: 156,
			Weight: 48,
		},
		{
			Name:   "Jan",
			Age:    29,
			Height: 188,
			Weight: 90,
		},
		{
			Name:   "Tay",
			Age:    33,
			Height: 190,
			Weight: 87,
		},
	}
	test := team.Oldest()
	// "{Tay 33 190 87}"
	wantedResult := team[3]
	if test != wantedResult {

		t.Errorf("team.Oldest() = %v; want %v", test, wantedResult)
	}
}

func TestPeopleGreet(t *testing.T) {
	team := People{
		{Name: "Anna",
			Age:    17,
			Height: 160,
			Weight: 49,
		},
		{
			Name:   "Luna",
			Age:    30,
			Height: 156,
			Weight: 48,
		},
		{
			Name:   "Jan",
			Age:    29,
			Height: 188,
			Weight: 90,
		},
		{
			Name:   "Tay",
			Age:    33,
			Height: 190,
			Weight: 87,
		},
	}
	test := team.Greet()
	namesSlice := []string{}
	// wantedResult := []string{}
	for _, firstName := range team {
		namesSlice = append(namesSlice, firstName.Name)
	}

	for _, value := range namesSlice {
		for _, familyName := range namesSlice {
			if value == familyName {
				continue
			}
		}
	}

	// 	numbers := []int{0, 42, 8}
	// 	fmt.Println(slices.Equal(numbers, []int{0, 42, 8}))
	// 	fmt.Println(slices.Equal(numbers, []int{10}))
	// }

	// numbers := []string{"A", "B", "CD"}
	// fmt.Println(slices.Equal(numbers, []string{"A", "B", "CD"}))
	// fmt.Println(slices.Equal(numbers, []string{"C"}))

	wantedResult := []string{
		"Hi Anna from Luna!",
		"Hi Anna from Jan!",
		"Hi Anna from Tay!",
		"Hi Luna from Anna!",
		"Hi Luna from Jan!",
		"Hi Luna from Tacdy!",
		"Hi Jan from Anna!",
		"Hi Jan from Luna!",
		"Hi Jan from Tay!",
		"Hi Tay from Anna!",
		"Hi Tay from Luna!",
		"Hi Tay from Jan!",
	}
	// if test != wantedResult {

	if slices.Equal(test, wantedResult) {

		t.Errorf("team.Greet() = %v; want %v", test, wantedResult)
	}
}
