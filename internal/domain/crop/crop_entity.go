package crop

type Crop struct {
	id          int64
	cropType    CropType
	isIrrigated bool
	isInsured   bool
}

func NewCrop(cropType string, isIrrigated bool, isInsured bool) (*Crop, error) {
	newCropType, err := NewCropType(cropType)
	if err != nil {
		return nil, err
	}

	return &Crop{
		cropType:    newCropType,
		isIrrigated: isIrrigated,
		isInsured:   isInsured,
	}, nil
}

func NewCropWithId(cropId int64, cropType string, isIrrigated bool, isInsured bool) (*Crop, error) {
	newCropType, err := NewCropType(cropType)
	if err != nil {
		return nil, err
	}

	return &Crop{
		id:          cropId,
		cropType:    newCropType,
		isIrrigated: isIrrigated,
		isInsured:   isInsured,
	}, nil

}
