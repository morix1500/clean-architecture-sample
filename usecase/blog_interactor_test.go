package usecase

import (
	"github.com/morix1500/clean-architecture-sample/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = domain.Blog{
	Id:      1,
	Title:   "Test",
	Content: "Hoge",
}

type dummyBlogRepository struct{}

func (dummyBlogRepository) Insert(domain.Blog) error {
	return nil
}

func (dummyBlogRepository) Select(id int32) (domain.Blog, error) {
	return testData, nil
}

func TestInsert(t *testing.T) {
	blogInteractor := BlogInteractor{BlogRepository: dummyBlogRepository{}}

	err := blogInteractor.Add(testData)
	assert.Equal(t, err, nil)
}

func TestSelect(t *testing.T) {
	blogInteractor := BlogInteractor{BlogRepository: dummyBlogRepository{}}

	var id int32 = 1
	b, err := blogInteractor.BlogById(id)
	expect := testData
	assert.Equal(t, b, expect)
	assert.Equal(t, err, nil)
}
