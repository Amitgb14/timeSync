package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestExecCommand(t *testing.T) {
	actualDate := ExecCommand()

	expectedDate := fmt.Sprintf("%s", time.Now())

	if actualDate != expectedDate {
		t.Fatalf("Expected %s but got %s", expectedDate, actualDate)
	}
}
