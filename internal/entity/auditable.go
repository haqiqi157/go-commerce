package entity

import (
	"gorm.io/gorm"
	"time"
)

type Auditable struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func NewAuditable() Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
}

func UpdateAuditable() Auditable {
	return Auditable{
		UpdatedAt: time.Now(),
	}
}

//func DeleteAuditable() Auditable {
//	return Auditable{
//		DeletedAt: gorm.DeletedAt{}(),
//	}
//}
