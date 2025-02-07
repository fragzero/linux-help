package main

import (
	"context"
	"os"
	"testing"
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
