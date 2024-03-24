package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

const hotelFacilitiesPath = "hotel.txt"
const roomFacilitiesPath = "room.txt"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}

func InsertHotelFacilities(fileName string) error {
	open, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)

	content, err := io.ReadAll(open)
	if err != nil {
		return err
	}
	var FacilitiesType []HotelFacilityType

	err = json.Unmarshal(content, &FacilitiesType)
	if err != nil {
		return err
	}

	for _, item := range FacilitiesType {
		if err := globaldb.Create(&item).Error; err != nil {
			return err
		}
		for _, facility := range item.Categories {
			facility.Type = item.Id
			if err := globaldb.Create(&facility).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertRoomFacilities(fileName string) error {
	open, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)

	content, err := io.ReadAll(open)
	if err != nil {
		return err
	}
	var FacilitiesType []RoomFacilityType

	err = json.Unmarshal(content, &FacilitiesType)
	if err != nil {
		return err
	}

	for _, item := range FacilitiesType {
		if err := globaldb.Create(&item).Error; err != nil {
			return err
		}
		for _, facility := range item.Categories {
			facility.Type = item.Id
			if err := globaldb.Create(&facility).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var globaldb *gorm.DB

func main() {
	mySqlConnString := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(mySqlConnString), &gorm.Config{})
	if err != nil {
		log.Fatal(db, err)
	}
	db = db.Debug()
	globaldb = db
	fmt.Println(InsertHotelFacilities(hotelFacilitiesPath))
	fmt.Println(InsertRoomFacilities(roomFacilitiesPath))
}
