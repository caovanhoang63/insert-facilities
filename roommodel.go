package main

type RoomFacilityType struct {
	Facility   `json:",inline"`
	Categories []RoomFacility `json:"categories" gorm:"-"`
}

func (RoomFacilityType) TableName() string {
	return "room_facility_types"
}

type RoomFacility FacilityDetail

func (RoomFacility) TableName() string {
	return "room_facilities"
}
