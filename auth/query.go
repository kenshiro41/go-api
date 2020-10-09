package auth

import (
	"github.com/kenshiro41/go_app/gql/models"
	"github.com/kenshiro41/go_app/utils"
)

func TokenCheck(userName string, token string) (*models.Message, error) {
	user, err := DecodeUser(token)

	if (err != nil) || (user.UserName != userName) {
		return utils.FailedMessage, nil
	}

	return utils.SuccessMessage, nil

}
