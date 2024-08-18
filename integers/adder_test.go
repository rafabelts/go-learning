package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
    operation := Add(2, 2)
    expected := 4
    assertCorrectMessage(t, operation, expected)
}

// Ejemplo testeable
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}

func assertCorrectMessage(t testing.TB, operation, expected int) {
    t.Helper()
    if(operation != expected) {
        t.Errorf("Got %d, expected %d", operation, expected)
    }
}
