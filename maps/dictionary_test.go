package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("known word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		assertErrors(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "word"
		definition := "this is a word"

		err := dictionary.Add(word, definition)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		word := "test"
		definition := "this is a new test"
		err := dictionary.Add(word, definition)

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, "this is a test")

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {

		word := "word"
		definition := "this is a word"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})
	t.Run("new word", func(t *testing.T) {
		word := "word"
		definition := "this is a word"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	want := definition
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, want)
}
