package model

import "time"

type TableOptions interface {
	TableOptions() string
}

var Migrations []interface{}

func Register(m interface{}) {
	Migrations = append(Migrations, m)
}

type BaseModel struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	//UpdatedAt time.Time  `json:"updated_at,omitempty"`
	//CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time `json:"-"`
	//DeletedAt *time.Time `gorm:"index" json:"-"`
}
