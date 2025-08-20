package dto

type ReviewRequest struct {
	CommitMessage string `json:"commitMessage"`
	Diff          string `json:"diff"`
}

type ReviewResponse struct {
	Feedback         string   `json:"summary"`
	Suggest string   `json:"suggestedMessage"`		
	Score     int      `json:"hygieneScore"`
}