package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	if geminiApiKey == "" {
		log.Fatal("Please set GEMINI_API_KEY environment variable")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./linux-help \"<prompt>\"")
		fmt.Println("Use quotes around the prompt.")
		fmt.Println("Example: ./linux-help \"how to find big files in a directory\"")
		fmt.Println("SHIFT+CTRL+V to paste the command.")
		os.Exit(1)
	}
	prompt := os.Args[1]

	// models: gemini-1.5-flash, gemini-1.5-flash-8b, gemini-1.0-pro
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.5)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("you are a linux terminal assistant, you can only answer with one linux commands, no markdown or comments or explanations"),
		},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	if err := copyToClipboard(resp); err != nil {
		log.Printf("Failed to copy to clipboard: %v", err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}

func copyToClipboard(resp *genai.GenerateContentResponse) error {
	var result string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%v", part)
			}
		}
	}
	return clipboard.WriteAll(result)
}
