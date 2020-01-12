package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestTableWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name     string
				LastName string
			}{"Chris", "Evans"},
			[]string{"Chris", "Evans"},
		},
		{
			"Struct with int and string field",
			struct {
				Name string
				Age  int
			}{"Chris", 35},
			[]string{"Chris"},
		}, {
			"nested struct",
			Person{
				"Chris",
				Profile{32, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"nested poiter to struct",
			&Person{
				"Chris",
				Profile{32, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slice input",
			[]Person{
				Person{"Chris", Profile{32, "London"}},
				Person{"Andy", Profile{32, "New York"}},
			},
			[]string{"Chris", "London", "Andy", "New York"},
		}, {
			"slice input",
			[2]Person{
				Person{"Chris", Profile{32, "London"}},
				Person{"Andy", Profile{32, "New York"}},
			},
			[]string{"Chris", "London", "Andy", "New York"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
