package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
    
    t.Run("Collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}
        got := Sum(numbers)
        want := 15
        assertCorrectMessage(t, got, want, numbers)
    })

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

    /*
    Go no deja usar comparadores con slices, por lo que se utiliza
    reflect.DeepEqual q es util para ver si dos variables son lo mismo
    reflect.DeepEqual no es type safety
    */

    assertSlicesCorrectMessage(t, got, want)

}

func TestSumAllTails(t *testing.T) {

    t.Run("make the sum of tails", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}

        assertSlicesCorrectMessage(t, got, want)
    })

    t.Run("safely run empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{0, 9})
        want := []int{0, 9}

        assertSlicesCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want int, numbers []int){
    t.Helper()
    
    if got != want {
        t.Errorf("got %d wanted %d given, %v", got, want, numbers)
    }
}


func assertSlicesCorrectMessage(t testing.TB, got, want []int){
    t.Helper()
    
    if !slices.Equal(got, want) {
        t.Errorf("got %v wanted %v", got, want)
    }
}
