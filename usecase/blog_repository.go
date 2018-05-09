package usecase

import (
	"github.com/morix1500/clean-architecture-sample/domain"
)

type BlogRepository interface {
	Insert(domain.Blog) error
	Select(int32) (domain.Blog, error)
}
