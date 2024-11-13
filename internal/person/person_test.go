package person

import (
	"slices"
	"testing"
)

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
	tests := []struct {
		p    Person
		want float64
	}{
		{
			p: Person{
				Name:   "Anna",
				Age:    17,
				Height: 160,
				Weight: 49,
			},
			want: 19.140625,
		},
		{
			p: Person{
				Name:   "Luna",
				Age:    30,
				Height: 156,
				Weight: 48,
			},
			want: 19.723865877712033,
		},
		{
			p: Person{
				Name:   "Jan",
				Age:    29,
				Height: 188,
				Weight: 90,
			},
			want: 25.464010864644635},
	}

	for _, test := range tests {
		got := test.p.BMI()
		if got != test.want {
			t.Errorf("BMI(%v) = %v; want %v", test.p, got, test.want)
		}
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		p    Person
		want bool
	}{
		{
			p: Person{
				Name:   "Anna",
				Age:    17,
				Height: 160,
				Weight: 49,
			},
			want: false,
		},
		{
			p: Person{
				Name:   "Luna",
				Age:    30,
				Height: 156,
				Weight: 48,
			},
			want: true,
		},
		{
			p: Person{
				Name:   "Jan",
				Age:    29,
				Height: 188,
				Weight: 90,
			},
			want: true,
		},
		{
			p: Person{
				Name:   "Tay",
				Age:    18,
				Height: 190,
				Weight: 87,
			},
			want: true},
	}

	for _, test := range tests {
		got := test.p.IsAdult()
		if got != test.want {
			t.Errorf("IsAdult(%v) = %v, want %v", test.p, got, test.want)
		}
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

	if slices.Equal(test, wantedResult) {
		t.Errorf("team.Greet() = %v; want %v", test, wantedResult)
	}
}
