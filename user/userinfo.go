package user

import (
	"github.com/kenshiro41/go_app/auth"
	sqls "github.com/kenshiro41/go_app/db/sql"
	"github.com/kenshiro41/go_app/gql/models"
	"github.com/kenshiro41/go_app/upload"
)

func UserInfo(userName string) (*models.User, error) {
	user := &models.User{}

	if err := db.Table("users").Where("user_name = ?", userName).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func TweetByUser(token string, userName string, current int) ([]*models.TweetData, error) {
	tweetsData := []*models.TweetData{}

	user := models.User{}
	db.Table("users").Where("user_name = ?", userName).Scan(&user)

	if err := db.Raw(sqls.TweetByUser, user.ID, current).Scan(&tweetsData).Error; err != nil {
		return nil, err
	}

	return tweetsData, nil
}

func UpdateProfile(input models.UpdateProfile, token string) (*models.Token, error) {
	decodeUser, err := auth.DecodeUser(token)
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

	t, err := auth.GenToken(decodeUser.ID, input.UserName)
	t.User = user

	return t, nil
}
