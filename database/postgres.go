package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sirArthurDayne/rest-ws/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)", post.Id, post.PostContent, post.UserId)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	userRow, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = userRow.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	user := models.User{}

	for userRow.Next() {
		if err = userRow.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = userRow.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	userRow, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1;", email)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = userRow.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	user := models.User{}

	for userRow.Next() {
		if err = userRow.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = userRow.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	postRow, err := repo.db.QueryContext(ctx, "SELECT id, post_content, user_id, created_at FROM posts WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = postRow.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	post := models.Post{}

	for postRow.Next() {
		if err = postRow.Scan(&post.Id, &post.PostContent, &post.UserId, &post.CreatedAt); err == nil {
			return &post, nil
		}
	}

	if err = postRow.Err(); err != nil {
		return nil, err
	}
	return &post, nil
}

func (repo *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 and user_id = $3",
		post.PostContent, post.Id, post.UserId)
	return err
}

func (repo *PostgresRepository) DeletePost(ctx context.Context, id, userId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE from posts WHERE id = $1 and user_id = $2", id, userId)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
