package repository

import "github.com/nathanhannon/bed-and-breakfast/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
