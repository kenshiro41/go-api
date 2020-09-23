package tweet

import (
	"time"

	"github.com/kenshiro41/go_app/auth"
	mydb "github.com/kenshiro41/go_app/db"
	"github.com/kenshiro41/go_app/gql/models"
	"github.com/kenshiro41/go_app/upload"
	"github.com/kenshiro41/go_app/utils"
)

var db = mydb.DB

func NewTweet(input models.NewTweet) (*models.TweetData, error) {
	now := time.Now()
	tweetName := utils.RandomString()

	decodeUser, err := auth.DecodeUser(input.Token)
	if err != nil {
		return nil, err
	}

	tweet := &models.Tweet{
		TweetName: tweetName,
		Text:      input.Text,
		UserID:    decodeUser.ID,
		CreatedAt: now,
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	if err := db.Table("tweets").Create(&tweet).Error; err != nil {
		return nil, err
	}

	imgs := []*models.Img{}

	for _, img := range input.Imgs {
		filename, err := upload.UploadImage(img)
		if err != nil {
			return nil, err
		}
		img := &models.Img{
			TweetID:   tweet.ID,
			ImgURL:    *filename,
			CreatedAt: now,
			UpdatedAt: nil,
			DeletedAt: nil,
		}
		if err := db.Table("imgs").Create(&img).Error; err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}

	tweetData := &models.TweetData{
		TweetName:    tweetName,
		Text:         input.Text,
		CreatedAt:    tweet.CreatedAt,
		UserID:       decodeUser.ID,
		UserName:     decodeUser.UserName,
		Nickname:     decodeUser.UserName,
		UserImg:      nil,
		ImgCount:     0,
		CommentCount: 0,
		FavCount:     0,
		IsFavorite:   0,
	}

	return tweetData, nil
}

func UpdateFavs(input models.UpdateFavorite) (*models.Message, error) {
	fav := &models.Favorite{}

	decodeUser, err := auth.DecodeUser(input.Token)
	if err != nil {
		return utils.FailedMessage, nil
	}

	if input.IsFavorite { // お気に入り登録
		fav = &models.Favorite{
			UserID:  decodeUser.ID,
			TweetID: input.TweetID,
		}
		if err := db.Table("favorites").Where("user_id = ? AND tweet_id = ?", decodeUser.ID, input.TweetID).Create(&fav).Error; err != nil {
			return utils.FailedMessage, nil
		}

	} else { //お気に入り解除
		if err := db.Table("favorites").Where("user_id = ? AND tweet_id = ?", decodeUser.ID, input.TweetID).Delete(&fav).Error; err != nil {
			return utils.FailedMessage, nil
		}
	}

	return utils.SuccessMessage, nil
}

func RemoveTweet(tweetID int, token string) (*models.Message, error) {
	decodeUser, err := auth.DecodeUser(token)
	if err != nil {
		return utils.FailedMessage, err
	}

	now := time.Now()
	tweet := &models.Tweet{
		DeletedAt: &now,
	}

	if err := db.Table("tweets").Where("id = ? AND user_id = ?", tweetID, decodeUser.ID).Update(&tweet).Error; err != nil {
		return utils.FailedMessage, err
	}

	return utils.SuccessMessage, nil
}

func UpdateTweet(input models.UpdateTweet) (*models.TweetData, error) {
	d := &models.TweetData{}

	decodeUser, err := auth.DecodeUser(input.Token)
	if err != nil {
		return nil, err
	}

	if err := db.Table("tweets").Where("id = ? AND user_id = ?", input.TweetID, decodeUser.ID).Update("text", input.Text).Error; err != nil {
		return nil, err
	}

	raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
		users.id AS user_id, users.user_name, users.nickname, users.user_img,
		(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
		(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count,
		(SELECT CAST(COUNT(1) AS BIT) FROM favorites WHERE favorites.tweet_id = tweets.id AND favorites.user_id = ?) AS isFavorite
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		WHERE tweets.deleted_at IS NULL
		AND tweets.id = ? AND users.id = ?
		LIMIT 1`

	if err := db.Raw(raw, decodeUser.ID, input.TweetID, decodeUser.ID).Scan(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}
