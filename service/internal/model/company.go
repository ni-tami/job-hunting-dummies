package model

import "time"

type Company struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	User        *User      `json:"user"                 gorm:"-;foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	CompanyName string     `json:"company_name"`
	Website     string     `json:"website"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
