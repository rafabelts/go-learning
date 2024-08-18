package main
import "testing" // libreria de testeo

func TestHello(t *testing.T) {

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Rafa", "")
        want := "Hello, Rafa"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello in spanish", func(t *testing.T) {
        got := Hello("Rafa", "Spanish")
        want := "Hola, Rafa"
        assertCorrectMessage(t, got, want)
    })
    
    t.Run("say 'Hello world' when an empty string is supplied", func(t * testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if (got != want) {
        t.Errorf("got %q, want %q", got, want)
    }
}
