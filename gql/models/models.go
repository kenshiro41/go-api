// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Comment struct {
	Comment   string     `json:"comment"`
	TweetID   int        `json:"tweet_id"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CommentInfo struct {
	Comment   string    `json:"comment"`
	TweetID   int       `json:"tweet_id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Nickname  string    `json:"nickname"`
	UserImg   *string   `json:"user_img"`
	CreatedAt time.Time `json:"created_at"`
}

type Favorite struct {
	TweetID   int       `json:"tweet_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Follow struct {
	FollowingID int        `json:"following_id"`
	FollowedID  int        `json:"followed_id"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type FollowCounts struct {
	Followings int `json:"followings"`
	Followers  int `json:"followers"`
}

type Img struct {
	ImgURL    string     `json:"img_url"`
	TweetID   int        `json:"tweet_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Message struct {
	Success bool `json:"success"`
}

type Token struct {
	Token string `json:"token"`
	Iat   int    `json:"iat"`
	Exp   int    `json:"exp"`
	User  *User  `json:"user"`
}

type Tweet struct {
	ID        int        `json:"id"`
	TweetName string     `json:"tweet_name"`
	Text      string     `json:"text"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type TweetData struct {
	ID           int        `json:"id"`
	TweetName    string     `json:"tweet_name"`
	Text         string     `json:"text"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UserID       int        `json:"user_id"`
	UserName     string     `json:"user_name"`
	Nickname     string     `json:"nickname"`
	UserImg      *string    `json:"user_img"`
	ImgCount     int        `json:"ImgCount"`
	CommentCount int        `json:"CommentCount"`
	FavCount     int        `json:"FavCount"`
	IsFavorite   int        `json:"isFavorite"`
}

type User struct {
	ID        int        `json:"id"`
	UserName  string     `json:"user_name"`
	Nickname  string     `json:"nickname"`
	Password  string     `json:"password"`
	UserImg   *string    `json:"user_img"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserInfo struct {
	User   *User        `json:"user"`
	Tweets []*TweetData `json:"tweets"`
}

type AddComment struct {
	Token   string `json:"token"`
	TweetID int    `json:"tweet_id"`
	Comment string `json:"comment"`
}

type FollowingInfo struct {
	IsFollowing int     `json:"isFollowing"`
	Following   []*User `json:"following"`
	Followed    []*User `json:"followed"`
}

type NewTweet struct {
	Token string   `json:"token"`
	Text  string   `json:"text"`
	Imgs  []string `json:"imgs"`
}

type UpdateFavorite struct {
	TweetID    int    `json:"tweet_id"`
	Token      string `json:"token"`
	IsFavorite bool   `json:"isFavorite"`
}

type UpdateFollow struct {
	Token       string `json:"token"`
	FollowedID  int    `json:"followed_id"`
	FolowStatus bool   `json:"folowStatus"`
}

type UpdateProfile struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Img      string `json:"img"`
}

type UpdateTweet struct {
	Token   string `json:"token"`
	Text    string `json:"text"`
	TweetID int    `json:"tweet_id"`
}