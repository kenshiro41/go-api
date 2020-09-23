package user

import (
	"time"

	"github.com/kenshiro41/go_app/utils"

	"github.com/kenshiro41/go_app/gql/models"
)

func Signup(userName string, password string) (*models.Message, error) {
	now := time.Now()

	user := &models.User{
		UserName:  userName,
		Password:  password,
		Nickname:  userName,
		UserImg:   nil,
		CreatedAt: now,
		UpdatedAt: nil,
	}

	if err := db.Table("users").Create(&user).Error; err != nil {
		return utils.FailedMessage, err
	}

	return utils.SuccessMessage, nil
}
