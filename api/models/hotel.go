package models

type CreateHotel struct {
	HotelName     string         `json:"hotel_name" default:"Silk Road"`
	Description   string         `json:"description" default:"in affordable prices"`
	Rating        float64        `json:"rating" default:"4.6"`
	ContactNumber string         `json:"contact_number" default:"+(99891)-234-56-78"`
	LicenceUrl    string         `json:"licence_url" default:"https://creativecommons.org/licenses/by/1.2/"`
	WebsiteUrl    string         `json:"website_url" default:"https://creativecommons.org/licenses/by/1.3/"`
	Images        []*CreateImage `json:"images"`
	Location      CreateLocation `json:"location"`
}

type HotelModel struct {
	HotelId       string        `json:"hotel_id"`
	OwnerId       string        `json:"owner_id"`
	HotelName     string        `json:"hotel_name"`
	Description   string        `json:"description"`
	Rating        float32       `json:"rating"`
	ContactNumber string        `json:"contact_number"`
	LicenceUrl    string        `json:"licence_url"`
	WebsiteUrl    string        `json:"website_url"`
	Images        []*ImageModel `json:"images"`
	Location      LocationModel `json:"location"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

type ListHotelsModel struct {
	Hotels []*HotelModel `json:"hotels"`
	Count  uint64        `json:"count"`
}

type UpdateHotel struct {
	HotelName     string         `json:"hotel_name" default:"updated hotel name"`
	Description   string         `json:"description" default:"updated description"`
	Rating        float64        `json:"rating" default:"5.0"`
	ContactNumber string         `json:"contact_number" default:"updated contact number"`
	LicenceUrl    string         `json:"licence_url" default:"updated licence url"`
	WebsiteUrl    string         `json:"website_url" default:"updated website url"`
	Location      UpdateLocation `json:"location"`
}