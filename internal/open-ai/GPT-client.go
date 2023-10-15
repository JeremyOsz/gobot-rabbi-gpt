package GPTClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GPTRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func SendGPTRequest(prompt string, maxTokens int) (string, error) {
	// Create the request body
	requestBody := GPTRequest{
		Prompt:    prompt,
		MaxTokens: maxTokens,
	}
	jsonBody, err := json.Marshal(requestBody)
	// print jsonBody to stdout
	fmt.Println(string(jsonBody))
	if err != nil {
		return "Error Marchalling JSON", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "Error making request", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI_API_KEY"))

	// Set the prompt

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error connecting to client", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != 200 {
		// Print the response status code to stdout
		fmt.Println(resp.StatusCode)
		// Read the response body error message
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err != nil {
			return "Error Parsing error message", err
		}
		return "Error: Response status code is not 200", nil
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error Parsing response", err
	}

	// Unmarshal the response body
	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		return "", err
	}

	// Check if the response is empty
	if len(gptResponse.Choices) == 0 {
		// Print gptResponse to stdout in
		fmt.Println(gptResponse)
		return "Error: Response is empty", nil
	}

	// Return the generated text
	return gptResponse.Choices[0].Text, nil
}
