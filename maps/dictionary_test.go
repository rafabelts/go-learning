package main

import (
	"testing"
)

func TestDictionary(t *testing.T) {
    dictionary := Dictionary{"test": "this is a test"}

    t.Run("known word", func(t *testing.T){
        got, _ := dictionary.Search("test")
        want := "this is a test"

        assertCorrectMessage(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T){
        _, err := dictionary.Search("unknown")
        want := ErrNotFound

        if err == nil {
            t.Fatal("expected to get an error")
        }

        assertError(t, err, want)
    })

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
       
        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
	})

    t.Run("alreay added word", func(t *testing.T) {
        word := "test"
        definition := "this is just a tes"

        dictionary := Dictionary{word: definition}

        err := dictionary.Add(word, definition)

        assertError(t, err, ErrWordExists)

        assertDefinition(t, dictionary, word, definition)
    })
}

func TestUpdate(t *testing.T) {
    t.Run("updating existent word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        newDef := "new"

        err := dictionary.Update(word, newDef)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, newDef)
    })

    t.Run("updating non existent word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        newDef := "new"

        err := dictionary.Update(word, newDef)
        assertError(t, err, ErrWordNotExists)
    })
}

func TestDelete(t *testing.T) {
    t.Run("delete existant word", func(t *testing.T) {
        word := "test"
        dictionary := Dictionary{word: "this is going to be deleted"}

        deleteErr := dictionary.Delete(word)
        assertError(t, deleteErr, nil)

        _, err := dictionary.Search(word)
        assertError(t, err, ErrNotFound)
    })
    
    t.Run("delete non existant word", func(t *testing.T) {
        dictionary := Dictionary{}

        err := dictionary.Delete("test")

        assertError(t, err, ErrWordNotExists)
    })

}

func assertError(t testing.TB, got, want error) {
    t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()
    
    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word: ", err)
    }

    assertCorrectMessage(t, got, definition)
}
