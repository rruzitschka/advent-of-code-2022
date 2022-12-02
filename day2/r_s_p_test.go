package main

import "testing"

func TestMap(t *testing.T) {
	// this tests the mapping of the strategy input to R, S, P
	// test cases:
	// input A=Rock B=Paper C=Scissor  X=Rock Y=Paper, Z=Scissors
	// (A, X) -> "R"
	// (B, Y) -> "P"
	// (C, Z) -> "S"

	got := mapTool("A")
	want := "R"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

	got = mapTool("X")
	want = "R"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

	got = mapTool("B")
	want = "P"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

	got = mapTool("Y")
	want = "P"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

	got = mapTool("C")
	want = "S"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

	got = mapTool("Z")
	want = "S"

	if got != want {
		t.Errorf("Score got %s, want %s", got, want)
	}

}

func TestWinRound(t *testing.T) {
	// this test the basic game functionality and the result is win score
	// if player 1 wins the score is 0, for a draw it is 3, if player two wins the score is 6
	// input is "R", "P", "S"
	// (R, P) -> 6
	// (S, P) -> 0
	// (P, P) -> 3
	// (S, R) -> 6

	got := winRound("R", "P")
	want := 6
	if got != want {
		t.Errorf("Score got %d, want %d", got, want)
	}

	got = winRound("S", "P")
	want = 0
	if got != want {
		t.Errorf("Score got %d, want %d", got, want)
	}

	got = winRound("P", "P")
	want = 3
	if got != want {
		t.Errorf("Score got %d, want %d", got, want)
	}

	got = winRound("S", "R")
	want = 6
	if got != want {
		t.Errorf("Score got %d, want %d", got, want)
	}

}

func TestWinRoundScore(t *testing.T) {
	// input A=Rock B=Paper C=Scissors
	// strategy X=Rock Y=Paper, Z=Scissors
	// points: win =6 draw = 3 lose = 0
	// additional points: Rock +1, paper +2, scissors +3
	// Input: Player 1 (A,B,C), Player 2 (X,Y,Z)
	// test cases :
	// (A, Y) -> player 2 wins, score 8
	// (B, X) -> player 1 wins, score 1
	// (C, Z) -> draw, score 6

	got := RoundScore(mapTool("A"), mapTool("Y"))
	want := 8

	if got != want {
		t.Errorf("Score got %d, want %d", got, want)
	}

}

func TestPLayer2ToolNewStrategy(t *testing.T) {
	// input A=Rock B=Paper C=Scissors
	// strategy X=Lose Y=Draw, Z=Win
	// test cases:
	// (A,X) -> S
	// (B,Y) -> P
	// (A,Y) -> R
	// (C,Z) -> R
	// (B,Z) -> S

	got := player2ToolNewStrategy(mapTool("A"), "X")
	want := "S"

	if got != want {
		t.Errorf("PLayer2 should use  %s, but used  %s", want, got)
	}

	got = player2ToolNewStrategy(mapTool("B"), "Y")
	want = "P"

	if got != want {
		t.Errorf("PLayer2 should use  %s, but used  %s", want, got)
	}

	got = player2ToolNewStrategy(mapTool("B"), "Z")
	want = "S"

	if got != want {
		t.Errorf("PLayer2 should use  %s, but used  %s", want, got)
	}

}
