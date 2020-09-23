package user

import (
	"github.com/kenshiro41/go_app/auth"
	"github.com/kenshiro41/go_app/gql/models"
)

func Login(userName string, password string) (*models.Token, error) {
	user := &models.User{}
	if err := db.Table("users").Where("user_name = ? AND password = ?", userName, password).Scan(&user).Error; err != nil {
		return nil, err
	}

	token, err := auth.GenToken(user.ID, user.UserName)
	if err != nil {
		return nil, err
	}

	token.User = &models.User{
		ID:       user.ID,
		UserName: user.UserName,
		Nickname: user.Nickname,
	}

	return token, nil
}
