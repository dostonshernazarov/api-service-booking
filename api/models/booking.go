package models

import "github.com/google/uuid"

type CreateBookingReq struct {
	HraId          string `json:"hra_id"`
	WillArrive     string `json:"will_arrive"`
	WillLeave      string `json:"will_leave"`
	NumberOfPeople int64  `json:"number_of_people"`
	IsCanceled     bool   `json:"is_canceled"`
	Reason         string `json:"reason"`
}

type UpdateBookingReq struct {
	Id             uuid.UUID `json:"id"`
	HraId          string    `json:"hra_id"`
	WillArrive     string    `json:"will_arrive"`
	WillLeave      string    `json:"will_leave"`
	NumberOfPeople int64     `json:"number_of_people"`
	IsCanceled     bool      `json:"is_canceled"`
	Reason         string    `json:"reason"`
}

type BookingRes struct {
	Id             uuid.UUID `json:"id"`
	UserId         string    `json:"user_id"`
	HraId          string    `json:"hra_id"`
	WillArrive     string    `json:"will_arrive"`
	WillLeave      string    `json:"will_leave"`
	NumberOfPeople int64     `json:"number_of_people"`
	IsCanceled     bool      `json:"is_canceled"`
	Reason         string    `json:"reason"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	DeletedAt      string    `json:"deleted_at"`
}

type IdReq struct {
	Id string `json:"id"`
}

type IdRes struct {
	Ids []*IdReq `json:"ids"`
}

type List struct {
	Bookings []*BookingRes `json:"bookings"`
}

type BookedUser struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BookedTime  string `json:"created_at"`
}
