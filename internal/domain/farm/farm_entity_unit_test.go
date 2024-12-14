package farm

import (
	"testing"
	"time"
)

func TestNewFarm(t *testing.T) {

	t.Run("Create Farm Successfully", func(t *testing.T) {
		farm, err := NewFarm("MyFarm", 100.5, "acre", "123 Farm Address")
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if farm.GetName() != "MyFarm" {
			t.Errorf("Expected farm name to be 'MyFarm', got: %v", farm.GetName())
		}

	})

	t.Run("Fail to Create Farm with Empty Name", func(t *testing.T) {
		_, err := NewFarm("", 100.5, "acre", "123 Farm Address")
		if err == nil {
			t.Fatal("Expected error for empty name, got none")
		}
	})

	t.Run("Fail to Create Farm with Invalid Unit", func(t *testing.T) {
		_, err := NewFarm("MyFarm", 100.5, "invalid_unit", "123 Farm Address")
		if err == nil {
			t.Fatal("Expected error for invalid unit, got none")
		}
	})

	t.Run("Fail to Create Farm with Empty Address", func(t *testing.T) {
		_, err := NewFarm("MyFarm", 100.5, "acre", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque cursus porta laoreet. Suspendisse nec mi ac velit pretium vehicula ac at velit. Integer turpis eros, sollicitudin id consectetur vel, dignissim eu velit. Aliquam gravida mattis leo, vitae sagittis odio viverra id. Duis non interdum mauris, ac dignissim ipsum. Etiam ornare eros libero, id placerat sem fringilla sed. Nulla in commodo nunc, ac fringilla dui. Duis at tempor est. Proin auctor mauris ipsum, scelerisque dignissim arcu pulvinar eu. Suspendisse eros lacus, porta vitae tincidunt eget, iaculis sit amet erat. Quisque at sapien velit. Aliquam nec velit a purus dignissim malesuada ac sollicitudin enim. Praesent ornare nunc augue, et condimentum nisi rhoncus id. Donec risus diam, sagittis non aliquam a, tristique ut.")
		if err == nil {
			t.Fatal("Expected error for empty address, got none")
		}
	})

	t.Run("Fail to Create Farm with Empty Address", func(t *testing.T) {
		_, err := NewFarm("MyFarm", 100.5, "acre", "")
		if err == nil {
			t.Fatal("Expected error for empty address, got none")
		}
	})
}

func TestNewFarmWithID(t *testing.T) {
	crops := []string{"corn", "wheat"}
	createdAt := time.Now()

	t.Run("Create Farm with ID Successfully", func(t *testing.T) {
		farm, err := NewFarmWithID("abc", "MyFarm", 100.5, "acre", "123 Farm Address", crops, createdAt)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if farm.GetID() != "abc" {
			t.Errorf("Expected farm ID to be 1, got: %v", farm.GetID())
		}
	})

	t.Run("Fail to Create Farm with Invalid Data", func(t *testing.T) {
		_, err := NewFarmWithID("abc", "", 0, "invalid_unit", "", crops, createdAt)
		if err == nil {
			t.Fatal("Expected error for invalid data, got none")
		}
	})
}
