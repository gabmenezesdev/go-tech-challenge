package farm

import "testing"

func TestValidUnitOfMeasure(t *testing.T) {
	units := []string{
		UnitAcre,
		UnitHectare,
		UnitSquareKm,
	}

	for _, unit := range units {
		unitOfMeasure, err := newUnitOfMeasure(unit)

		if err != nil {
			t.Error("expect to have no errors")
		}

		if unitOfMeasure.Value() != unit {
			t.Errorf("expect value to be %v, but got %v", unit, unitOfMeasure.Value())
		}
	}
}

func TestInvalidUnitOfMeasure(t *testing.T) {
	_, err := newUnitOfMeasure("unit")

	if err == nil {
		t.Error("expect to have errors")
	}

	_, err = newUnitOfMeasure("")

	if err == nil {
		t.Error("expect to have errors")
	}
}
