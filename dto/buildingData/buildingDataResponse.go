package buildingDatadto

type BuildingDataResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name" form:"name"`
	PhoneNumber     string  `json:"phone_number" form:"phone_number"`
	Address         string  `json:"address" form:"address"`
	AllowedCapacity int     `json:"allowed_capacity" form:"allowed_capacity"`
	Longitude       float64 `json:"longitude" form:"longitude"`
	Latitude        float64 `json:"latitude" form:"latitude"`
}
