package usecase

import "gitcommit/domain/entity"

type IReviewUsecase interface {
	ReviewCommit(commit entity.Commit) (entity.Review, error)
}