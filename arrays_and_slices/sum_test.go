package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3}
	got := sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2,9}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0,9}

		checkSums(t, got, want)
	})

}
