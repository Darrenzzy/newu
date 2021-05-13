package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type RkResumeData struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	Show     int    `json:"show"` // 是否后台设置
	Errno    int    `json:"errno"`
}
