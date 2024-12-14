package farm

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
