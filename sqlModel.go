package main

import "time"

type Facility struct {
	SqlModel `json:",inline"`
	Name     string `json:"name" gorm:"column:name"`
	NameEn   string `json:"name_en" gorm:"column:name_en"`
	NameVn   string `json:"name_vn" gorm:"column:name_vn"`
}

type FacilityDetail struct {
	SqlModel      `json:",inline"`
	Name          string `json:"name" gorm:"column:name"`
	NameEn        string `json:"name_en" gorm:"column:name_en"`
	NameVn        string `json:"name_vn" gorm:"column:name_vn"`
	DescriptionEn string `json:"description_en" gorm:"column:description_en"`
	DescriptionVn string `json:"description_vn" gorm:"column:description_vn"`
	Type          int    `json:"type" gorm:"column:type"`
}

type SqlModel struct {
	Id        int        `json:"-" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
