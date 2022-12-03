package main

import "testing"

func TestSplitBagContent(t *testing.T) {

	testContent := "vJrwpWtwJgWrhcsFMMfFFhFp"

	got := splitBagContent(testContent)

	want := []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("got1 should be %s but was %s, got2 should be  %s, but was  %s \n", want[0], got[0], want[1], got[1])

	}

}

func TestFindSameCharacterInStringList(t *testing.T) {

	characterString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	stringList := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}

	got := findSameCharacterInStringList(stringList, characterString)
	want := "r"

	if got != want {
		t.Errorf(" List of three: Found %s but expected %s \n", got, want)
	}

	stringList = []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}

	got = findSameCharacterInStringList(stringList, characterString)
	want = "p"

	if got != want {
		t.Errorf(" List of two: Found %s but expected %s \n", got, want)
	}

}

func TestFindPriority(t *testing.T) {
	testCharacter := "p"

	got := findPriority(testCharacter)
	want := 16

	if got != want {
		t.Errorf("Got prio %d but expected Prio %d \n", got, want)
	}
}
