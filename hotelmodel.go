package main

type HotelFacilityType struct {
	Facility   `json:",inline"`
	Categories []HotelFacility `json:"categories" gorm:"-"`
}

func (HotelFacilityType) TableName() string {
	return "hotel_facility_types"
}

type HotelFacility FacilityDetail

func (HotelFacility) TableName() string {
	return "hotel_facilities"
}
