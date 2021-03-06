package daos

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tombiers/estuary-backend/models"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Estuary?charset=utf8&parseTime=True")
	fmt.Println("booking controller init")

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}
}

// GetAll returns all entries
func GetAll() []models.Booking {
	bookings := []models.Booking{}
	db.Find(&bookings)
	return bookings
}

// GetById returns entry with the given id
func GetById(id int) models.Booking {
	var booking models.Booking
	db.First(&booking, id)
	return booking
}

// Create creates a new entry
func Create(booking models.Booking) {
	db.Create(&booking)
}

// Update the entry with the given id to the given data
func Update(id int, booking models.Booking) models.Booking {
	var dbBooking models.Booking
	db.Where("ID = ?", id).First(&dbBooking)
	dbBooking.User = booking.User
	dbBooking.Members = booking.Members
	db.Save(&dbBooking)
	return dbBooking
}

// Delete the entry with the given id
func Delete(id int) {
	var booking models.Booking
	db.First(&booking, id)
	db.Delete(&booking)
}
