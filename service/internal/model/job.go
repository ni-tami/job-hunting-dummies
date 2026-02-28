package model

import "time"

type Job struct {
	ID           int64      `json:"id"`
	Company      *Company   `json:"company" gorm:"-;foreignKey:company_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	CompanyID    int64      `json:"company_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Requirements []string   `json:"requirements" gorm:"serializer:json"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}
