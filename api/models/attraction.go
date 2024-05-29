package models

import "time"

type Attraction struct {
	AttractionId   string    `json:"attraction_id"`
	OwnerId        string    `json:"owner_id"`
	AttractionName string    `json:"attraction_name"`
	Description    string    `json:"description"`
	Rating         float64   `json:"rating"`
	ContactNumber  string    `json:"contact_number"`
	LicenceUrl     string    `json:"licence_url"`
	WebsiteUrl     string    `json:"website_url"`
	Images         []*Image  `json:"images"`
	Location       Location  `json:"location"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type Image struct {
	ImageId         string    `json:"image_id"`
	EstablishmentId string    `json:"establishment_id"`
	ImageUrl        string    `json:"image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type Location struct {
	LocationId      string    `json:"location_id"`
	EstablishmentId string    `json:"establishment_id"`
	Address         string    `json:"address"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Country         string    `json:"country"`
	City            string    `json:"city"`
	StateProvince   string    `json:"state_province"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type CreateAttraction struct {
	AttractionName string         `json:"attraction_name" default:"Anhor Park"`
	Description    string         `json:"description" default:"available for all ages"`
	Rating         float64        `json:"rating" default:"4.3"`
	ContactNumber  string         `json:"contact_number" default:"+(99891)-234-56-78"`
	LicenceUrl     string         `json:"licence_url" default:"https://creativecommons.org/licenses/by/4.0/"`
	WebsiteUrl     string         `json:"website_url" default:"https://creativecommons.org/licenses/by/4.1/"`
	Images         []*CreateImage `json:"images"`
	Location       CreateLocation `json:"location"`
}

type CreateImage struct {
	ImageUrl string `json:"image_url" default:"www.photo/images/141"`
}

type CreateLocation struct {
	Address       string  `json:"address" default:"87G9+V57, Shaykhontohur Street 28, Tashkent, Toshkent Shahri"`
	Latitude      float64 `json:"latitude" default:"40.7128"`
	Longitude     float64 `json:"longitude" default:"74.0060"`
	Country       string  `json:"country" default:"Uzbekistan"`
	City          string  `json:"city" default:"Tashkent"`
	StateProvince string  `json:"state_province" default:"Shaykhontohur"`
}

type AttractionModel struct {
	AttractionId   string        `json:"attraction_id"`
	OwnerId        string        `json:"owner_id"`
	AttractionName string        `json:"attraction_name"`
	Description    string        `json:"description"`
	Rating         float32       `json:"rating"`
	ContactNumber  string        `json:"contact_number"`
	LicenceUrl     string        `json:"licence_url"`
	WebsiteUrl     string        `json:"website_url"`
	Images         []*ImageModel `json:"images"`
	Location       LocationModel `json:"location"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
}

type ImageModel struct {
	ImageId         string `json:"image_id"`
	EstablishmentId string `json:"establishment_id"`
	ImageUrl        string `json:"image_url"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type LocationModel struct {
	LocationId      string  `json:"location_id"`
	EstablishmentId string  `json:"establishment_id"`
	Address         string  `json:"address"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Country         string  `json:"country"`
	City            string  `json:"city"`
	StateProvince   string  `json:"state_province"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ListAttractionModel struct {
	Attractions []*AttractionModel `json:"attractions"`
	Count       uint64             `json:"count"`
}

type UpdateAttraction struct {
	AttractionName string         `json:"attraction_name" default:"updated attraction name"`
	Description    string         `json:"description" default:"updated description"`
	Rating         float64        `json:"rating" default:"5.0"`
	ContactNumber  string         `json:"contact_number" default:"updated contact number"`
	LicenceUrl     string         `json:"licence_url" default:"updated licence url"`
	WebsiteUrl     string         `json:"website_url" default:"updated website url"`
	Location       UpdateLocation `json:"location"`
}

type UpdateLocation struct {
	Address       string  `json:"address" default:"updated address"`
	Latitude      float64 `json:"latitude" default:"1.1"`
	Longitude     float64 `json:"longitude" default:"1.1"`
	Country       string  `json:"country" default:"updated country"`
	City          string  `json:"city" default:"updated city"`
	StateProvince string  `json:"state_province" default:"updated state or province"`
}

type DeleteResponse struct {
	Success bool `json:"success"`
}

type FieldValuesByLocation struct {
	Country string `json:"country"`
	City string `json:"city"`
	Province string `json:"province"`
}

type FindByName struct {
	Name string `json:"name"`
}