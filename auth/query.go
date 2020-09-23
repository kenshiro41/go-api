package auth

import (
	"github.com/kenshiro41/go_app/gql/models"
)

func TokenCheck(userName string, token string) (*models.Message, error) {
	falseMessage := &models.Message{
		Success: false,
	}
	successMessage := &models.Message{
		Success: true,
	}

	user, err := DecodeUser(token)

	if (err != nil) || (user.UserName != userName) {
		return falseMessage, nil
	}

	return successMessage, nil

}
