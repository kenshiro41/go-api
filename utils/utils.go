package utils

import (
	mydb "github.com/kenshiro41/go_app/db"

	"github.com/kenshiro41/go_app/gql/models"
)

func MapTweetID(tweets []*models.Tweet, userID *int) (*[]*models.TweetData, error) {
	db := mydb.DB

	tweetsData := []*models.TweetData{}
	raw := `SELECT *,
		(SELECT COUNT(*) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(*) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		LEFT OUTER JOIN imgs ON imgs.tweet_id = tweets.id
		WHERE tweets.deleted_at IS NULL
		ORDER BY tweets.created_at DESC`

	if err := db.Raw(raw, &userID).Scan(&tweetsData).Error; err != nil {
		return nil, err
	}

	// for _, t := range tweets {

	// 	user := &models.User{}
	// 	if err := db.Table("users").Where("id = ?", t.UserID).Scan(&user).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	tweet := &models.Tweet{}
	// 	if err := db.Table("tweets").Where("id = ?", t.ID).Scan(&tweet).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	var commentCount int

	// 	if err := db.Table("comments").Where("tweet_id = ?", t.ID).Count(&commentCount).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	imgs := []*models.Img{}
	// 	if err := db.Table("imgs").Where("tweet_id = ?", t.ID).Scan(&imgs).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	var favCount int
	// 	if err := db.Table("favorites").Where("tweet_id = ?", t.ID).Count(&favCount).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	isFav := false
	// 	if userID != nil {
	// 		favorite := &models.Favorite{}
	// 		if err := db.Table("favorites").Where("tweet_id = ? AND user_id = ?", t.ID, userID).Scan(&favorite).Error; err != nil {
	// 			isFav = false
	// 		}
	// 		if *userID == favorite.UserID {
	// 			isFav = true
	// 		}
	// 	}

	// 	data := &models.TweetData{
	// 		User:         user,
	// 		Tweet:        tweet,
	// 		CommentCount: commentCount,
	// 		Imgs:         imgs,
	// 		FavCount:     favCount,
	// 		IsFavorite:   isFav,
	// 	}

	// 	tweetsData = append(tweetsData, data)
	// }

	return &tweetsData, nil
}
