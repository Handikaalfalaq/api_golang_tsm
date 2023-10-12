package models

type BuildingData struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	PhoneNumber     string  `json:"phone_number"`
	Address         string  `json:"address"`
	AllowedCapacity int     `json:"allowed_capacity"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
}
