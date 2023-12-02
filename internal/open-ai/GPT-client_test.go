package GPTClient

import "testing"

// TestSendRequest tests the SendRequest function
func TestSendRequest(t *testing.T) {
	//  Try run AskGPT
	response, err := AskGPT("Hello, GPT!", 50)
	if err != nil {
		t.Errorf("AskGPT() error = %v", err)
	}

	//  assert that the response is not empty
	if len(response) == 0 {
		t.Errorf("Response is empty")
	}

}

//  Test asking what this week's Torah portion is
func TestPortion(t *testing.T) {
	//  Try run AskGPT
	response, err := AskGPT("Can you summarise Parshat Bereshit", 500)
	if err != nil {
		t.Errorf("AskGPT() error = %v", err)
	}

	//  assert that the response is not empty
	if len(response) == 0 {
		t.Errorf("Response is empty")
	}
}
