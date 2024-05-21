package models

import "github.com/google/uuid"

type CreateBookingReq struct {	
	HraId string `json:"hra_id"`
	WillArrive string `json:"will_arrive"`
	WillLeave string `json:"will_leave"`
	NumberOfPeople int64 `json:"number_of_people"`
	IsCanceled bool `json:"is_canceled"`
	Reason string `json:"reason"`
}

type UpdateBookingReq struct {
	Id uuid.UUID `json:"id"`
	UserId string `json:"user_id"`
	HraId string `json:"hra_id"`
	WillArrive string `json:"will_arrive"`
	WillLeave string `json:"will_leave"`
	NumberOfPeople int64 `json:"number_of_people"`
	IsCanceled bool `json:"is_canceled"`
	Reason string `json:"reason"`
}

type BookingRes struct {
	Id uuid.UUID `json:"id"`
	UserId string `json:"user_id"`
	HraId string `json:"hra_id"`
	WillArrive string `json:"will_arrive"`
	WillLeave string `json:"will_leave"`
	NumberOfPeople int64 `json:"number_of_people"`
	IsCanceled bool `json:"is_canceled"`
	Reason string `json:"reason"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type GetAllByUIdReq struct {
	Limit uint64 `json:"limit"`
	Page uint64 `json:"page"`
	UserId string `json:"user_id"`
}

type GetAllByHRAIdReq struct {
	Limit uint64 `json:"limit"`
	Page uint64 `json:"page"`
	HraId string `json:"hra_id"`
}

type UserId struct {
	Id string `json:"id"`
}

type GetAllByUIdRes struct {
    Bookings []*BookingRes `json:"bookings"`
}

type GetAllByHRAIdRes struct {
    UserId []*UserId `json:"user_id"`
}

type List struct {
    Bookings []*BookingRes `json:"bookings"`
}

