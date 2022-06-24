package store

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Post struct {
	ID         int
	Title      string `binding:"required,min=3,max=50"`
	Content    string `binding:"required,min=5,max=5000"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	UserID     int `json:"-"`
}

func AddPost(user *User, post *Post) error {
	post.UserID = user.ID
	_, err := db.Model(post).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("Error inserting new post")
	}
	return err
}

func FetchUserPosts(user *User) ([]Post, error) {
	posts := []Post{}
	err := db.Model(&posts).Where("user_id = ?", user.ID).Select()
	// err := db.Model(user).
	// 	Relation("Posts", func(q *orm.Query) (*orm.Query, error) {
	// 		return q.Order("id ASC"), nil
	// 	}).Select()
	if err != nil {
		log.Error().Err(err).Msg("Error fetching user's posts")
	}
	return posts, err
}

func FetchPost(id int) (*Post, error) {
	post := new(Post)
	post.ID = id
	err := db.Model(post).WherePK().Select()
	if err != nil {
		log.Error().Err(err).Msg("Error fetching post")
		return nil, err
	}
	return post, nil
}

func UpdatePost(post *Post) error {
	_, err := db.Model(post).WherePK().UpdateNotZero()
	if err != nil {
		log.Error().Err(err).Msg("Error updating post")
	}
	return err
}

func DeletePost(post *Post) error {
	_, err := db.Model(post).WherePK().Delete()
	if err != nil {
		log.Error().Err(err).Msg("Error deleting post")
	}
	return err
}
