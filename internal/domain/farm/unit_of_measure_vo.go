package farm

import "errors"

type UnitOfMeasure struct {
	value string
}

const (
	UnitAcre     = "acre"
	UnitHectare  = "hectare"
	UnitSquareKm = "square_km"
)

func newUnitOfMeasure(unit string) (UnitOfMeasure, error) {
	if len(unit) == 0 {
		return UnitOfMeasure{}, errors.New("Unit is empty")
	}
	switch unit {
	case UnitAcre, UnitHectare, UnitSquareKm:
		return UnitOfMeasure{value: unit}, nil
	default:
		return UnitOfMeasure{}, errors.New("Unit is invalid")
	}
}

func (u *UnitOfMeasure) Value() string {
	return u.value
}
