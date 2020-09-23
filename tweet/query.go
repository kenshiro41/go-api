package tweet

import (
	"github.com/kenshiro41/go_app/utils"

	"github.com/kenshiro41/go_app/auth"
	"github.com/kenshiro41/go_app/gql/models"
)

func AllTweet(token *string, current int) ([]*models.TweetData, error) {
	d := []*models.TweetData{}

	if token != nil { // Adminユーザー用
		decodeUser, err := auth.DecodeUser(*token)
		if err != nil {
			return nil, err
		}

		raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
			users.id AS user_id, users.user_name, users.nickname, users.user_img,
			(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
			(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
			(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count,
			(SELECT CAST(COUNT(1) AS BIT) FROM favorites WHERE favorites.tweet_id = tweets.id AND favorites.user_id = ?) AS is_favorite
			FROM tweets
			INNER JOIN users ON tweets.user_id = users.id
			WHERE tweets.deleted_at IS NULL
			ORDER BY tweets.created_at DESC`
		// OFFSET ? LIMIT 10`

		if err := db.Raw(raw, &decodeUser.ID).Scan(&d).Error; err != nil {
			return nil, err
		}

	} else { // ランディングページ用
		raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
		users.id AS user_id, users.user_name, users.nickname, users.user_img,
		(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
		(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		WHERE tweets.deleted_at IS NULL
		ORDER BY tweets.created_at DESC`

		if err := db.Raw(raw).Scan(&d).Error; err != nil {
			return nil, err
		}
	}

	return d, nil
}

func TweetByID(tweetID int) (*models.TweetData, error) {
	d := &models.TweetData{}

	raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
		users.id AS user_id, users.user_name, users.nickname, users.user_img,
		(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
		(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		WHERE tweets.deleted_at IS NULL
		AND tweets.id = ?
		LIMIT 1`

	if err := db.Raw(raw, tweetID).Scan(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func Search(text string) ([]*models.TweetData, error) {
	search := "%" + text + "%"

	d := []*models.TweetData{}

	raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
		users.id AS user_id, users.user_name, users.nickname, users.user_img,
		(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
		(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		WHERE tweets.deleted_at IS NULL
		AND tweets.text LIKE ? 
		OR users.user_name LIKE ?
		OR users.nickname LIKE ?
		ORDER BY tweets.created_at DESC`
	if err := db.Raw(raw, search, search, search).Scan(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func CheckCanEdit(token string, tweetID int) (*models.Message, error) {
	decodUser, err := auth.DecodeUser(token)
	if err != nil {
		return utils.FailedMessage, err
	}

	tweet := &models.Tweet{}
	if err := db.Table("tweets").Select("id, user_id, created_at, deleted_at").Where("id = ? AND user_id = ?", tweetID, decodUser.ID).Scan(&tweet).Error; err != nil {
		return utils.FailedMessage, err
	}

	return utils.SuccessMessage, nil
}

func ImageByID(tweetID int) ([]*models.Img, error) {
	img := []*models.Img{}
	if err := db.Table("imgs").Where("tweet_id = ?", tweetID).Scan(&img).Error; err != nil {
		return nil, err
	}

	return img, nil
}
