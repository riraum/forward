package person

import "testing"

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
	test := anna.BMI()
	wantedResult := 19.140625
	if test != wantedResult {

		t.Errorf("anna.BMI() = %v; want %v", test, wantedResult)
	}
}

func TestIsAdult(t *testing.T) {
	anna := Person{
		Name:   "Anna",
		Age:    17,
		Height: 160,
		Weight: 49,
	}
	test := anna.IsAdult()
	wantedResult := false
	if test != wantedResult {

		t.Errorf("anna.IsAdult() = %v; want %v", test, wantedResult)
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

// func TestPeopleGreet(t *testing.T) {
// 	team := People{
// 		{Name: "Anna",
// 			Age:    17,
// 			Height: 160,
// 			Weight: 49,
// 		},
// 		{
// 			Name:   "Luna",
// 			Age:    30,
// 			Height: 156,
// 			Weight: 48,
// 		},
// 		{
// 			Name:   "Jan",
// 			Age:    29,
// 			Height: 188,
// 			Weight: 90,
// 		},
// 		{
// 			Name:   "Tay",
// 			Age:    33,
// 			Height: 190,
// 			Weight: 87,
// 		},
// 	}
// 	test := team.Greet()
// 	wantedResult := []string{"Hi Anna from Luna!",
// 		"Hi Anna from Jan!",
// 		"Hi Anna from Tay!",
// 		"Hi Luna from Anna!",
// 		"Hi Luna from Jan!",
// 		"Hi Luna from Tacdy!",
// 		"Hi Jan from Anna!",
// 		"Hi Jan from Luna!",
// 		"Hi Jan from Tay!",
// 		"Hi Tay from Anna!",
// 		"Hi Tay from Luna!",
// 		"Hi Tay from Jan!"}
// 	if test != wantedResult {

// 		t.Errorf("team.Greet() = %v; want %v", test, wantedResult)
// 	}
// }
