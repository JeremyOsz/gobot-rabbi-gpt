package GPTClient

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func getAPIKey() string {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the API key from the .env file
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPEN_AI_API_KEY environment variable is not set")
	}
	return apiKey
}

func AskGPT(prompt string, maxTokens int) (string, error) {
	// Send tthe request
	response, err := SendRequest(prompt, maxTokens)
	if err != nil {
		return "Response empty", err
	}

	answer := response.Choices[0].Message.Content

	fmt.Println("Reb GPT: ", answer)
	return answer, nil

}

func SendRequest(prompt string, maxTokens int) (Response, error) {
	apiKey := getAPIKey()
	client := resty.New()

	// Send the request to GPT-3
	response, err := sendPostRequest(client, apiKey, prompt, maxTokens)
	if err != nil {
		return Response{}, err
	}

	// Unmarshal the response to read the JSON as a struct
	gptResponse, err := unmarshalResponse(response)
	if err != nil {
		return Response{}, err
	}

	if len(gptResponse.Choices) == 0 {
		fmt.Println(gptResponse)
		return Response{}, nil
	}

	return gptResponse, nil
}

func sendPostRequest(client *resty.Client, apiKey, prompt string, maxTokens int) (*resty.Response, error) {
	return client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []interface{}{map[string]interface{}{
				"role":    "system",
				"content": prompt,
			}},
			"max_tokens": maxTokens,
		}).
		Post(apiEndpoint)
}

func unmarshalResponse(response *resty.Response) (Response, error) {
	var gptResponse Response
	err := json.Unmarshal(response.Body(), &gptResponse)
	return gptResponse, err
}
