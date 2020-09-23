package utils

import "github.com/kenshiro41/go_app/gql/models"

var FailedMessage = &models.Message{
	Success: false,
}
var SuccessMessage = &models.Message{
	Success: true,
}
