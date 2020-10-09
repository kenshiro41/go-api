package sqls

var NotAdminData = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	WHERE tweets.deleted_at IS NULL
	ORDER BY tweets.created_at DESC
	LIMIT 10 OFFSET ?`

var AdminData = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count,
	(SELECT CAST(COUNT(1) AS BIT) FROM favorites WHERE favorites.tweet_id = tweets.id AND favorites.user_id = ?) AS is_favorite
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	WHERE tweets.deleted_at IS NULL
	ORDER BY tweets.created_at DESC
	LIMIT 10 OFFSET ?`

var Search = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	AND tweets.text LIKE ? 
	OR users.user_name LIKE ?
	OR users.nickname LIKE ?
	WHERE tweets.deleted_at IS NULL
	ORDER BY tweets.created_at DESC
	LIMIT 10 OFFSET ?`
var AdminSerach = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count,
	(SELECT CAST(COUNT(1) AS BIT) FROM favorites WHERE favorites.tweet_id = tweets.id AND favorites.user_id = ?) AS is_favorite
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	AND tweets.text LIKE ? 
	OR users.user_name LIKE ?
	OR users.nickname LIKE ?
	WHERE tweets.deleted_at IS NULL
	ORDER BY tweets.created_at DESC
	LIMIT 10 OFFSET ?`

var TweetByUser = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	WHERE tweets.deleted_at IS NULL
	AND users.id = ?
	ORDER BY tweets.created_at DESC
	LIMIT 10 OFFSET ?`

var UpdateTweet = `SELECT tweets.id, tweets.tweet_name, tweets.text, tweets.created_at,
	users.id AS user_id, users.user_name, users.nickname, users.user_img,
	(SELECT COUNT(img_url) FROM imgs WHERE imgs.tweet_id = tweets.id) AS img_count,
	(SELECT COUNT(id) FROM comments WHERE comments.tweet_id = tweets.id) AS comment_count,
	(SELECT COUNT(id) FROM favorites WHERE favorites.tweet_id = tweets.id) AS fav_count,
	(SELECT CAST(COUNT(1) AS BIT) FROM favorites WHERE favorites.tweet_id = tweets.id AND favorites.user_id = ?) AS isFavorite
	FROM tweets
	INNER JOIN users ON tweets.user_id = users.id
	WHERE tweets.deleted_at IS NULL
	AND tweets.id = ? AND users.id = ?
	LIMIT 1`
