package farm

import "errors"

type LandArea struct {
	value float64
}

func NewLandArea(landArea float64) (LandArea, error) {
	if landArea <= 0 {
		return LandArea{}, errors.New("landArea must be greater than 0")
	}

	return LandArea{value: landArea}, nil
}

func (l *LandArea) Value() float64 {
	return l.value
}
