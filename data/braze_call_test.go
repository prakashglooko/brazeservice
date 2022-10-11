package data

import "testing"

func TestMissingRequiredFields(t *testing.T) {
	bc := &BrazeCall{}

	err := bc.Validate()

	if err == nil {
		t.Fatal("Nil fields allowed")
	}
}

func TestAvailablRequiredFields(t *testing.T) {
	bc := &BrazeCall{"braze_data", "response", "questionnaire"}

	err := bc.Validate()

	if err != nil {
		t.Fatal("Non Nil fields errored")
	}
}
