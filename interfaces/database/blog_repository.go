package database

import (
	"github.com/morix1500/clean-architecture-sample/domain"
)

type BlogRepository struct {
	SqlHandler
}

func (repo *BlogRepository) Insert(b domain.Blog) error {
	_, err := repo.Execute(
		"INSERT INTO blog (id, title, content) VALUES (?, ?, ?)", b.Id, b.Title, b.Content,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *BlogRepository) Select(blogId int32) (domain.Blog, error) {
	row, err := repo.Query("SELECT id, title, content FROM blog WHERE id = ?", blogId)
	defer row.Close()
	if err != nil {
		return domain.Blog{}, err
	}

	var id int32
	var title, content string

	row.Next()
	if err = row.Scan(&id, &title, &content); err != nil {
		return domain.Blog{}, err
	}
	return domain.Blog{
		Id: id,
		Title: title,
		Content: content,
	}, nil
}
