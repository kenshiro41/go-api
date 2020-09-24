package user

import (
	"github.com/kenshiro41/go_app/auth"
	"github.com/kenshiro41/go_app/gql/models"
	"github.com/kenshiro41/go_app/upload"
)

func UserInfo(userName string) (*models.UserInfo, error) {
	user := &models.User{}

	if err := db.Table("users").Where("user_name = ?", userName).Scan(&user).Error; err != nil {
		return nil, err
	}

	tweetsData := []*models.TweetData{}
	raw := `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
		users.id AS user_id, users.user_name, users.nickname, users.user_img,
		(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
		(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
		(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS favorite_count
		FROM tweets
		INNER JOIN users ON tweets.user_id = users.id
		WHERE tweets.deleted_at IS NULL
		AND users.id = ?
		ORDER BY tweets.created_at DESC`

	if err := db.Raw(raw, user.ID).Scan(&tweetsData).Error; err != nil {
		return nil, err
	}

	userInfo := &models.UserInfo{
		User:   user,
		Tweets: tweetsData,
	}

	return userInfo, nil
}

func UpdateProfile(input models.UpdateProfile) (*models.Token, error) {
	decodeUser, err := auth.DecodeUser(input.Token)
	if err != nil {
		return nil, err
	}

	var userImg *string

	if input.Img != "" {
		userImg, err = upload.UploadIcon(input.Img)
		if err != nil {
			return nil, err
		}
	}

	user := &models.User{
		UserName: input.UserName,
		Nickname: input.Nickname,
		UserImg:  userImg,
	}

	if err := db.Table("users").Where("id = ?", &decodeUser.ID).Update(&user).Error; err != nil {
		return nil, err
	}

	token, err := auth.GenToken(decodeUser.ID, input.UserName)
	token.User = user

	return token, nil
}
