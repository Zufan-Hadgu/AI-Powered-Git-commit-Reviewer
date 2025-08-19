package repository

import "gitcommit/domain/entity"

type IReviewRepository interface {
	ReviewCommit(commit entity.Commit) (*entity.Review, error)
}