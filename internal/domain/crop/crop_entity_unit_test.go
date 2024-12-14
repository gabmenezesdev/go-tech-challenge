package crop

import (
	"testing"
)

func TestNewCrop(t *testing.T) {
	t.Run("Valid crop creation", func(t *testing.T) {
		crop, err := NewCrop(Rice, true, false)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if crop.cropType.value != Rice {
			t.Errorf("Expected crop type to be %v, but got %v", Rice, crop.cropType.value)
		}

		if !crop.isIrrigated {
			t.Error("Expected crop to be irrigated, but it's not")
		}

		if crop.isInsured {
			t.Error("Expected crop to not be insured, but it is")
		}
	})

	t.Run("Invalid crop type", func(t *testing.T) {
		_, err := NewCrop("INVALID_CROP", true, true)
		if err == nil {
			t.Error("Expected error for invalid crop type, but got none")
		}
	})
}

func TestNewCropWithId(t *testing.T) {
	t.Run("Valid crop creation with ID", func(t *testing.T) {
		cropId := int64(42)
		crop, err := NewCropWithId(cropId, Corn, false, true)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if crop.id != cropId {
			t.Errorf("Expected crop ID to be %v, but got %v", cropId, crop.id)
		}

		if crop.cropType.value != Corn {
			t.Errorf("Expected crop type to be %v, but got %v", Corn, crop.cropType.value)
		}

		if crop.isIrrigated {
			t.Error("Expected crop to not be irrigated, but it is")
		}

		if !crop.isInsured {
			t.Error("Expected crop to be insured, but it's not")
		}
	})

	t.Run("Invalid crop type with ID", func(t *testing.T) {
		_, err := NewCropWithId(1, "INVALID_CROP", false, false)
		if err == nil {
			t.Error("Expected error for invalid crop type, but got none")
		}
	})
}
