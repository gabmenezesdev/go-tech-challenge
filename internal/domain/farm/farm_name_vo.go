package farm

import "errors"

type FarmName struct {
	value string
}

func NewFarmName(name string) (FarmName, error) {
	if len(name) == 0 {
		return FarmName{}, errors.New("name must exist")
	}
	if len(name) > 100 {
		return FarmName{}, errors.New("name must be less than 100 characters")
	}
	return FarmName{
		value: name,
	}, nil
}

func (n FarmName) Value() string {
	return n.value
}
