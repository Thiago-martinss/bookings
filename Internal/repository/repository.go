package repository

import "github.com/thiago-martinss/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}