package comment

import (
	"github.com/kenshiro41/go_app/auth"
	mydb "github.com/kenshiro41/go_app/db"
	"github.com/kenshiro41/go_app/gql/models"
)

var db = mydb.DB

func AddComment(input models.AddComment, token string) (*models.Comment, error) {
	decodeUser, err := auth.DecodeUser(token)
	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Comment: input.Comment,
		TweetID: input.TweetID,
		UserID:  decodeUser.ID,
	}

	if err := db.Table("comments").Create(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
