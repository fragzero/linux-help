package main

import (
	"context"
	"os"
	"testing"

	"github.com/atotto/clipboard"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func TestRetrieveGEMINI_API_KEY(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "test_key")
	defer os.Unsetenv("GEMINI_API_KEY")

	key := os.Getenv("GEMINI_API_KEY")
	if key != "test_key" {
		t.Errorf("Expected GEMINI_API_KEY to be 'test_key', got '%s'", key)
	}
}

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("test_key"))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
}

func TestCommandLineArgumentParsing(t *testing.T) {
	os.Args = []string{"cmd"}
	if len(os.Args) < 2 {
		t.Log("Usage: ./linux-help \"<prompt>\"")
		t.Log("Use quotes around the prompt.")
		t.Log("Example: ./linux-help \"how to find big files in a directory\"")
		t.Log("SHIFT+CTRL+V to paste the command.")
	}
}

func TestGenerateContent(t *testing.T) {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, option.WithAPIKey("test_key"))
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.5)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("you are a linux terminal assistant, you can only answer with one linux commands, no markdown or comments or explanations"),
		},
	}

	resp, err := model.GenerateContent(ctx, genai.Text("test prompt"))
	if err != nil {
		t.Fatalf("Failed to generate content: %v", err)
	}

	if len(resp.Candidates) == 0 {
		t.Error("Expected at least one candidate in response")
	}
}

func TestCopyToClipboard(t *testing.T) {
	resp := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []genai.Part{
						genai.Text("test content"),
					},
				},
			},
		},
	}

	err := copyToClipboard(resp)
	if err != nil {
		t.Fatalf("Failed to copy to clipboard: %v", err)
	}

	content, err := clipboard.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read from clipboard: %v", err)
	}

	if content != "test content" {
		t.Errorf("Expected clipboard content to be 'test content', got '%s'", content)
	}
}

func TestPrintResponse(t *testing.T) {
	resp := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []genai.Part{
						genai.Text("test content"),
					},
				},
			},
		},
	}

	printResponse(resp)
}
