package comment

import (
	"github.com/kenshiro41/go_app/gql/models"
)

func Comments(tweetID int) ([]*models.CommentInfo, error) {
	comments := []*models.CommentInfo{}

	if err := db.Select("comment, tweet_id, user_id, user_name, nickname, user_img, comments.created_at").Table("comments").Where("tweet_id = ?", tweetID).Joins("INNER JOIN users ON users.id = user_id").Scan(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
