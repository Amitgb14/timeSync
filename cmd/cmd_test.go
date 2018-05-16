package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestDateCommand(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	tt := time.Now()
	actualDate := DateCommand()
	expectedDate := fmt.Sprintf("%s", tt.Format(layout))

	if len(actualDate) == 0 {
		t.Fatalf("Expected %s but got %s", expectedDate, actualDate)
	}
}

func TestCheckNTPService(t *testing.T) {
	actualResponse := CheckNTPService()
	expectedResponse := true
	if actualResponse != expectedResponse {
		t.Fatalf("Expected %s but got %s", expectedResponse, actualResponse)
	}
}
