package crop

// Although the crop currently doesn't fully satisfy the second rule of an entity, I believe it has the potential to be a domain entity.

// Unique Identifier: The crop has a unique id.
// Independent Lifecycle: While not fully independent now, it can develop its own lifecycle with independent changes.
// Business Rules: It has or could gain specific business rules tied to its logic.

type Crop struct {
	id          int64
	cropType    CropType
	isIrrigated bool
	isInsured   bool
}

type CropDto struct {
	CropType    string `json:"crop_type"`
	IsIrrigated bool   `json:"is_irrigated"`
	IsInsured   bool   `json:"is_insured"`
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

func (c Crop) GetCropType() string {
	return c.cropType.Value()
}

func (c Crop) GetIsIrrigated() bool {
	return c.isIrrigated
}

func (c Crop) GetIsInsured() bool {
	return c.isInsured
}
