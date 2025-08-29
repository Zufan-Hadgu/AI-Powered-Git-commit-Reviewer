package controller

import (
	"gitcommit/domain/entity"
	"gitcommit/dto"
	"gitcommit/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	uc *usecase.ReviewUsecase
}

func NewReviewController(uc *usecase.ReviewUsecase) *ReviewController {
	return &ReviewController{uc: uc}
}

func (ctrl *ReviewController) Review(c *gin.Context) {
    var commitDto dto.ReviewRequest

    if err := c.ShouldBindJSON(&commitDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    commit := entity.Commit{
        Message: commitDto.CommitMessage,
        Diff:    commitDto.Diff,
    }

    reviewed, err := ctrl.uc.ReviewCommit(commit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    response := dto.ReviewResponse{
        Feedback: reviewed.Feedback,
        Suggest:  reviewed.Suggest,
        Score:    reviewed.Score,
    }

    c.JSON(http.StatusOK, response)
}
