package tweet

import (
	"fmt"

	sqls "github.com/kenshiro41/go_app/db/sql"

	"github.com/kenshiro41/go_app/utils"

	"github.com/kenshiro41/go_app/auth"
	"github.com/kenshiro41/go_app/gql/models"
)

func AllTweet(token string, current int) ([]*models.TweetData, error) {
	d := []*models.TweetData{}

	if token != "" { // Adminユーザー用
		decodeUser, err := auth.DecodeUser(token)
		if err != nil {
			return nil, err
		}

		if err := db.Raw(sqls.AdminData, &decodeUser.ID, current).Scan(&d).Error; err != nil {
			return nil, err
		}

	} else { // ランディングページ用
		if err := db.Raw(sqls.NotAdminData, current).Scan(&d).Error; err != nil {
			return nil, err
		}
	}
	fmt.Print(d[0].FavCount)
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

func Search(token string, text string, current int) ([]*models.TweetData, error) {
	search := "%" + text + "%"

	d := []*models.TweetData{}

	if token != "" { // Adminユーザー
		decodeUser, err := auth.DecodeUser(token)
		if err != nil {
			return nil, err
		}

		if err := db.Raw(sqls.AdminSerach, decodeUser.ID, search, search, search, current).Scan(&d).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Raw(sqls.Search, search, search, search, current).Scan(&d).Error; err != nil {
			return nil, err
		}
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
