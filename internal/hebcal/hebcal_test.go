package hebcal

import "testing"

func TestSendReque(t *testing.T) {
	//  Try run AskGPT
	_, err := GetWeeklyPortion("2023-12-02")
	if err != nil {
		t.Errorf("GetWeeklyPortion() error = %v", err)
	}

}
