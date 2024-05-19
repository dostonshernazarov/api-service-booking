package models

type CreateRestaurant struct {
	RestaurantName string         `json:"restaurant_name" default:"Kamolon Osh Markazi"`
	Description    string         `json:"description" default:"uzbek national cousine"`
	Rating         float64        `json:"rating" default:"4.1"`
	OpeningHours   string         `json:"opening_hours"  default:"06:00-22:00"`
	ContactNumber  string         `json:"contact_number" default:"+(99891)-234-56-78"`
	LicenceUrl     string         `json:"licence_url" default:"https://creativecommons.org/licenses/by/3.2/"`
	WebsiteUrl     string         `json:"website_url" default:"https://creativecommons.org/licenses/by/3.3/"`
	Images         []*CreateImage `json:"images"`
	Location       CreateLocation `json:"location"`
}

type RestaurantModel struct {
	RestaurantId   string        `json:"restaurant_id"`
	OwnerId        string        `json:"owner_id"`
	RestaurantName string        `json:"restaurant_name"`
	Description    string        `json:"description"`
	Rating         float32       `json:"rating"`
	OpeningHours   string        `json:"opening_hours"`
	ContactNumber  string        `json:"contact_number"`
	LicenceUrl     string        `json:"licence_url"`
	WebsiteUrl     string        `json:"website_url"`
	Images         []*ImageModel `json:"images"`
	Location       LocationModel `json:"location"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
}

type ListRestaurantsModel struct {
	Restaurants []*RestaurantModel `json:"restaurants"`
	Count       uint64             `json:"count"`
}

type UpdateRestaurant struct {
	RestaurantName string         `json:"restaurant_name" default:"updated restaurant name"`
	Description    string         `json:"description" default:"updated description"`
	Rating         float64        `json:"rating" default:"4.9"`
	OpeningHours   string         `json:"opening_hours" default:"09:00-00:00"`
	ContactNumber  string         `json:"contact_number" default:"updated contact number"`
	LicenceUrl     string         `json:"licence_url" default:"updated licence url"`
	WebsiteUrl     string         `json:"website_url" default:"updated website url"`
	Location       UpdateLocation `json:"location"`
}