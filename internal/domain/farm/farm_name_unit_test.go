package farm

import (
	"testing"
)

func TestValidName(t *testing.T) {
	farmName, err := NewFarmName("MyFarm")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if farmName.Value() != "MyFarm" {
		t.Errorf("Expected 'MyFarm', but got %v", farmName.Value())
	}
}

func TestInvalidName(t *testing.T) {
	_, err := NewFarmName("")
	if err == nil {
		t.Error("Expected error for empty name, but got none")
	}

	_, err = NewFarmName("This is a very long farm name that exceeds the limit. This is a very long farm name that exceeds the limit")
	if err == nil {
		t.Error("Expected error for name exceeding limit, but got none")
	}
}
