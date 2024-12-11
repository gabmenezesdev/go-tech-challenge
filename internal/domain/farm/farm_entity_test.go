package farm

import (
	"testing"
	"time"
)

func TestNewFarm(t *testing.T) {
	crops := []string{"corn", "wheat"}

	t.Run("Create Farm Successfully", func(t *testing.T) {
		farm, err := NewFarm("MyFarm", 100.5, "acre", "123 Farm Address", crops)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if farm.Name() != "MyFarm" {
			t.Errorf("Expected farm name to be 'MyFarm', got: %v", farm.Name())
		}

		if farm.GetID() != 0 {
			t.Errorf("Expected farm ID to be 0, got: %v", farm.GetID())
		}

		if farm.GetCreatedAt().IsZero() {
			t.Errorf("Expected createdAt to be set, but it is zero")
		}
	})

	t.Run("Fail to Create Farm with Empty Name", func(t *testing.T) {
		_, err := NewFarm("", 100.5, "acre", "123 Farm Address", crops)
		if err == nil {
			t.Fatal("Expected error for empty name, got none")
		}
	})

	t.Run("Fail to Create Farm with Invalid Unit", func(t *testing.T) {
		_, err := NewFarm("MyFarm", 100.5, "invalid_unit", "123 Farm Address", crops)
		if err == nil {
			t.Fatal("Expected error for invalid unit, got none")
		}
	})

	t.Run("Fail to Create Farm with Empty Address", func(t *testing.T) {
		_, err := NewFarm("MyFarm", 100.5, "acre", "", crops)
		if err == nil {
			t.Fatal("Expected error for empty address, got none")
		}
	})
}

func TestNewFarmWithID(t *testing.T) {
	crops := []string{"corn", "wheat"}
	createdAt := time.Now()

	t.Run("Create Farm with ID Successfully", func(t *testing.T) {
		farm, err := NewFarmWithID(1, "MyFarm", 100.5, "acre", "123 Farm Address", crops, createdAt)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if farm.GetID() != 1 {
			t.Errorf("Expected farm ID to be 1, got: %v", farm.GetID())
		}

		if !farm.GetCreatedAt().Equal(createdAt) {
			t.Errorf("Expected createdAt to be %v, got: %v", createdAt, farm.GetCreatedAt())
		}
	})

	t.Run("Fail to Create Farm with Invalid Data", func(t *testing.T) {
		_, err := NewFarmWithID(1, "", 0, "invalid_unit", "", crops, createdAt)
		if err == nil {
			t.Fatal("Expected error for invalid data, got none")
		}
	})
}
