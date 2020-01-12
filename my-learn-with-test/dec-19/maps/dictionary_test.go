package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")
	want := "this is just a test"
	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestAdd(t *testing.T) {
	t.Run("When key does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("When key exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")

		if err == nil {
			t.Fatal("Should throw an error")
		}
	})
}
