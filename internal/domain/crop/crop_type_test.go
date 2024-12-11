package crop

import "testing"

func TestValidCropType(t *testing.T) {
	validCrops := []string{
		Rice,
		Beans,
		Corn,
		Coffee,
		Soybeans,
	}

	for _, crop := range validCrops {
		cropType, err := NewCropType(crop)

		if err != nil {
			t.Errorf("Expected no error for crop %v, but got %v", crop, err)
		}

		if cropType.value != crop {
			t.Errorf("Expected value to be %v, but got %v", crop, cropType.value)
		}
	}
}

func TestInvalidCropType(t *testing.T) {
	invalidCrops := []string{
		"BANANA",
		"tomato",
		"",
		"123",
	}

	for _, crop := range invalidCrops {
		_, err := NewCropType(crop)

		if err == nil {
			t.Errorf("Expected error for invalid crop %v, but got none", crop)
		}
	}
}
