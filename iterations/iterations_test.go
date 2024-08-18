package iteration

import "testing"

func TestRepeat(t *testing.T){
    repeated := Repeat("a", 10)
    expected := "aaaaaaaaaa"
    assertCorrectMessage(t, repeated, expected)
} 

func BenchmarkRepeat(b *testing.B){
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
    t.Helper()

    if (repeated != expected) {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}
