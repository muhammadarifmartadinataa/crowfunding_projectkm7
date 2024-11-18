package utils

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func CallGeminiAPI() (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyCIWMbcw96aVutYTUhQM5XewWaV1HM016w"))
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("tolong buatkan saya deskripsi sebuah kampanye website bank sampah yang menarik dan jangan berikan simbol apapun"))
	if err != nil {
		return "", fmt.Errorf("failed to generate content from Gemini API: %v", err)
	}

	// Return the generated response
	return fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0]), nil
}
