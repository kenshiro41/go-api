package user

import (
	"fmt"
	"time"

	"github.com/kenshiro41/go_app/utils"

	"github.com/kenshiro41/go_app/auth"
	mydb "github.com/kenshiro41/go_app/db"
	"github.com/kenshiro41/go_app/gql/models"
)

var db = mydb.DB

func Follow(input models.UpdateFollow) (*models.Message, error) {
	now := time.Now()

	decodeUser, err := auth.DecodeUser(input.Token)
	if err != nil {
		return utils.FailedMessage, err
	}

	if input.FolowStatus {
		follow := &models.Follow{
			FollowingID: decodeUser.ID,
			FollowedID:  input.FollowedID,
			CreatedAt:   now,
		}
		if err := db.Table("followings").Create(&follow).Error; err != nil {
			return utils.FailedMessage, err
		}
	} else {
		fmt.Println(decodeUser.ID, input.FollowedID)
		raw := `DELETE FROM followings WHERE following_id = ? AND followed_id = ?`
		if err := db.Exec(raw, decodeUser.ID, input.FollowedID).Error; err != nil {
			return utils.FailedMessage, err
		}
	}

	return utils.SuccessMessage, nil
}

func CountFollowings(userID int) (*models.FollowCounts, error) {
	var followings int
	var followers int

	if err := db.Table("followings").Where("following_id = ?", userID).Count(&followings).Error; err != nil {
		return nil, err
	}
	if err := db.Table("followings").Where("followed_id = ?", userID).Count(&followers).Error; err != nil {
		return nil, err
	}

	counts := &models.FollowCounts{
		Followings: followings,
		Followers:  followers,
	}

	return counts, nil
}

func FollowInfo(UserName string, token *string) (*models.FollowingInfo, error) {
	followingUsers := []*models.User{}
	followedUsers := []*models.User{}
	followingQuery := "SELECT id, user_name, nickname, user_img FROM users WHERE id IN (SELECT followed_id FROM followings WHERE following_id = (SELECT id FROM users WHERE user_name = ?))"
	followedQuery := "SELECT id, user_name, nickname, user_img FROM users WHERE id IN (SELECT following_id FROM followings WHERE followed_id = (SELECT id FROM users WHERE user_name = ?))"

	if err := db.Raw(followingQuery, UserName).Scan(&followingUsers).Error; err != nil {
		return nil, err
	}
	if err := db.Raw(followedQuery, UserName).Scan(&followedUsers).Error; err != nil {
		return nil, err
	}

	isFollowing := 0
	if token != nil {
		decodeUser, err := auth.DecodeUser(*token)
		if err != nil {
			return nil, err
		}

		user := &models.User{}
		if err := db.Table("users").Where("user_name = ?", UserName).Scan(&user).Error; err != nil {
			return nil, err
		}

		raw := `SELECT (SELECT CAST(COUNT(0) AS BIT) FROM followings WHERE followings.following_id = ? 
		AND followings.followed_id = ?)AS isFollowing`
		if err := db.Raw(raw, decodeUser.ID, user.ID).Count(&isFollowing).Error; err != nil {
			return nil, err
		}
	}

	followingInfo := &models.FollowingInfo{
		IsFollowing: isFollowing,
		Following:   followingUsers,
		Followed:    followedUsers,
	}

	return followingInfo, nil
}
