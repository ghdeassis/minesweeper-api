package service

import "testing"

func TestPutFlag(t *testing.T) {
	flagRow := 1
	flagColumn := 1

	game := NewGame(9, 9, 4)
	game, err := PutFlag(game.ID, flagRow, flagColumn)
	if err != nil {
		t.Error(err)
		return
	}
	if !game.Cells[flagRow][flagColumn].IsFlagged {
		t.Errorf("PutFlag was incorrect, got false, want true. Row: %d, Column: %d", flagRow, flagColumn)
	}
}

func TestRemoveFlag(t *testing.T) {
	flagRow := 2
	flagColumn := 2

	game := NewGame(9, 9, 4)
	game, err := PutFlag(game.ID, flagRow, flagColumn)
	if err != nil {
		t.Error(err)
		return
	}
	game, err = RemoveFlag(game.ID, flagRow, flagColumn)
	if err != nil {
		t.Error(err)
		return
	}
	if game.Cells[flagRow][flagColumn].IsFlagged {
		t.Errorf("RemoveFlag was incorrect, got true, want false. Row: %d, Column: %d", flagRow, flagColumn)
	}
}
