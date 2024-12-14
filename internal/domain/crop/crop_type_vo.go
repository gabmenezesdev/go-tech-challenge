package crop

import "errors"

type CropType struct {
	value string
}

const (
	Rice     = "RICE"
	Beans    = "BEANS"
	Corn     = "CORN"
	Coffee   = "COFFEE"
	Soybeans = "SOYBEANS"
)

func NewCropType(crop string) (CropType, error) {
	switch crop {
	case Rice, Beans, Corn, Coffee, Soybeans:
		return CropType{value: crop}, nil
	default:
		return CropType{}, errors.New("Crop is invalid")
	}
}

func (c CropType) Value() string {
	return c.value
}
