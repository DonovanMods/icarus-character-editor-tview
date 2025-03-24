package data

import (
	"fmt"

	"github.com/donovanmods/icarus-player-data/character"
	"github.com/donovanmods/icarus-player-data/profile"
)

var (
	CharacterData *character.CharacterData
	ProfileData   *profile.ProfileData
)

func Read() error {
	var err error

	CharacterData, err = character.NewCharacterData()
	if err != nil {
		return fmt.Errorf("error reading character data: %w", err)
	}

	ProfileData, err = profile.NewProfileData()
	if err != nil {
		return fmt.Errorf("error reading profile data: %w", err)
	}

	return nil
}
