package usecase

import (
	"gitcommit/domain/entity"
	"gitcommit/domain/repository"
)

type ReviewUsecase struct {
	repo repository.IReviewRepository
}

func NewReviewUsecase(r repository.IReviewRepository) *ReviewUsecase {
	return &ReviewUsecase{repo: r}
}

func (u *ReviewUsecase) ReviewCommit(commit entity.Commit) (*entity.Review, error) {
	return u.repo.ReviewCommit(commit)

}