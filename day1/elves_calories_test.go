package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestSclice(t *testing.T) {

	testCaloriesData := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	t.Run("test the slicing", func(t *testing.T) {
		got := sliceCaloriesList(testCaloriesData)
		want := []int{6000, 4000, 11000, 24000, 10000}

		if !Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}

	})

}

func TestMaxCalories(t *testing.T) {

	testCalories := []int{6000, 4000, 11000, 24000, 10000}

	got := maxCalories(testCalories)
	want := 24000

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func TestTopThreeTotalCalories(t *testing.T) {

	testCalories := []int{6000, 4000, 11000, 24000, 10000}

	got := topThreeTotalCalories(testCalories)
	want := 45000

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
