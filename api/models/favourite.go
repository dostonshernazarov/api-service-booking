package models

import "time"

type Favourite struct {
	FavouriteId     string    `json:"favourite_id"`
	EstablishmentId string    `json:"establishment_id"`
	UserId          string    `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type CreateFavourite struct {
	EstablishmentId string `json:"establishment_id"`
	UserId          string `json:"user_id"`
}

type FavouriteModel struct {
	FavouriteId     string `json:"favourite_id"`
	EstablishmentId string `json:"establishment_id"`
	UserId          string `json:"user_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type RemoveResponse struct {
	Success bool `json:"success"`
}

type ListFavouritesModel struct {
	Favourites []*FavouriteModel `json:"favourites"`
}

