package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kenshiro41/go_app/auth"
	"github.com/kenshiro41/go_app/comment"
	"github.com/kenshiro41/go_app/gql/models"
	"github.com/kenshiro41/go_app/tweet"
	"github.com/kenshiro41/go_app/user"
	"github.com/kenshiro41/go_app/utils"
)

func (r *mutationResolver) Signup(ctx context.Context, userName string, password string) (*models.Message, error) {
	return user.Signup(userName, password)
}

func (r *mutationResolver) Login(ctx context.Context, userName string, password string) (*models.Token, error) {
	return user.Login(userName, password)
}

func (r *mutationResolver) CreateTweet(ctx context.Context, input models.NewTweet) (*models.TweetData, error) {
	t := ctx.Value(models.Token{}).(string)

	tweetData, err := tweet.NewTweet(input, t)
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	for _, observer := range addChannelObserver {
		observer <- tweetData
	}
	r.mu.Unlock()

	return tweetData, nil
}

func (r *mutationResolver) UpdateTweet(ctx context.Context, input models.UpdateTweet) (*models.TweetData, error) {
	t := ctx.Value(models.Token{}).(string)

	return tweet.UpdateTweet(input, t)
}

func (r *mutationResolver) DeleteTweet(ctx context.Context, tweetID int) (*models.Message, error) {
	t := ctx.Value(models.Token{}).(string)

	return tweet.RemoveTweet(tweetID, t)
}

func (r *mutationResolver) AddComment(ctx context.Context, input models.AddComment) (*models.Comment, error) {
	t := ctx.Value(models.Token{}).(string)

	return comment.AddComment(input, t)
}

func (r *mutationResolver) UpdateFavorite(ctx context.Context, input models.UpdateFavorite) (*models.Message, error) {
	t := ctx.Value(models.Token{}).(string)

	return tweet.UpdateFavs(input, t)
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input models.UpdateProfile) (*models.Token, error) {
	t := ctx.Value(models.Token{}).(string)

	return user.UpdateProfile(input, t)
}

func (r *mutationResolver) FollowUser(ctx context.Context, input models.UpdateFollow) (*models.Message, error) {
	t := ctx.Value(models.Token{}).(string)

	return user.Follow(input, t)
}

func (r *queryResolver) Tweets(ctx context.Context, current int) ([]*models.TweetData, error) {
	t := ctx.Value(models.Token{}).(string)

	return tweet.AllTweet(t, current)
}

func (r *queryResolver) TweetByID(ctx context.Context, tweetID int) (*models.TweetData, error) {
	return tweet.TweetByID(tweetID)
}

func (r *queryResolver) SearchText(ctx context.Context, text string, current int) ([]*models.TweetData, error) {
	t := ctx.Value(models.Token{}).(string)
	return tweet.Search(t, text, current)
}

func (r *queryResolver) ImageByID(ctx context.Context, tweetID int) ([]*models.Img, error) {
	return tweet.ImageByID(tweetID)
}

func (r *queryResolver) Comments(ctx context.Context, tweetID int) ([]*models.CommentInfo, error) {
	return comment.Comments(tweetID)
}

func (r *queryResolver) TokenCheck(ctx context.Context, userName string) (*models.Message, error) {
	t := ctx.Value(models.Token{}).(string)

	return auth.TokenCheck(userName, t)
}

func (r *queryResolver) EditCheck(ctx context.Context, tweetID int) (*models.Message, error) {
	t := ctx.Value(models.Token{}).(string)

	return tweet.CheckCanEdit(t, tweetID)
}

func (r *queryResolver) UserInfo(ctx context.Context, userName string) (*models.User, error) {
	return user.UserInfo(userName)
}

func (r *queryResolver) TweetByUser(ctx context.Context, userName string, current int) ([]*models.TweetData, error) {
	t := ctx.Value(models.Token{}).(string)
	return user.TweetByUser(t, userName, current)
}

func (r *queryResolver) FollowCount(ctx context.Context, userID int) (*models.FollowCounts, error) {
	return user.CountFollowings(userID)
}

func (r *queryResolver) FollowInfo(ctx context.Context, userName string) (*models.FollowingInfo, error) {
	t := ctx.Value(models.Token{}).(string)
	return user.FollowInfo(userName, &t)
}

func (r *subscriptionResolver) AddTweetChannel(ctx context.Context) (<-chan *models.TweetData, error) {
	id := utils.RandString(8)
	events := make(chan *models.TweetData, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(addChannelObserver, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	addChannelObserver[id] = events
	r.mu.Unlock()

	return events, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var addChannelObserver map[string]chan *models.TweetData

func init() {
	addChannelObserver = map[string]chan *models.TweetData{}
}
