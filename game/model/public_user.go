package model

import "github.com/olefen/note/game"

// PubCreateUser
func PubCreateUser() (ret *game.User, err error) {
	u := game.NewUser(nil)
	if err = ModelUser.Create(u).Error; err != nil {
		return
	}

	//
	ret = u
	return
}