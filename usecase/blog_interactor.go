package usecase

import (
	"github.com/morix1500/clean-architecture-sample/domain"
)

type BlogInteractor struct {
	BlogRepository BlogRepository
}

func (bi *BlogInteractor) Add(b domain.Blog) error {
	err := bi.BlogRepository.Insert(b)
	return err
}

func (bi *BlogInteractor) BlogById(id int32) (domain.Blog, error) {
	blog, err := bi.BlogRepository.Select(id)
	return blog, err
}
