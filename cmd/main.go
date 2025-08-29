package main

import (
	"gitcommit/controller"
	"gitcommit/infrastructure/ai"
	"gitcommit/repository"
	"gitcommit/usecase"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil{
		panic("GROK_API_KEY not in the env file")
	}
	// --- Load API Key ---
	apiKey := os.Getenv("GROK_API_KEY")
	if apiKey == "" {
		panic("GROK_API_KEY not set in environment variables")
	}

	// --- Infrastructure ---
	grokClient := ai.NewGrokClient(apiKey)

	// --- Repository ---
	reviewRepo := repository.NewReviewRepository(grokClient)

	// --- Usecase ---
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo)

	// --- Controller ---
	reviewController := controller.NewReviewController(reviewUsecase)

	// --- Router ---
	r := controller.SetupRouter(reviewController)

	// --- Run server ---
	r.Run(":8080")
}
