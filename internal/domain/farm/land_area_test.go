package farm

// if landArea <= 0 {
// 	return LandArea{}, errors.New("landArea must be greater than 0")
// }

// return LandArea{value: landArea}, nil

import (
	"testing"
)

func TestValidLandArea(t *testing.T) {
	landArea, err := NewLandArea(200)
	if err != nil {
		t.Error("expected no error")
	}

	if landArea.Value() != 200.00 {
		t.Error("expected no error")
	}
}

func TestInvalidLandArea(t *testing.T) {
	_, err := NewLandArea(0)
	if err == nil {
		t.Error("expected an invalid error, but get none")
	}
}
