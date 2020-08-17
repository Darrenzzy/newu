package models

import (
	orm "go-admin/global"
	"time"
)

type Member struct {
	Avatar           string    `gorm:"column:avatar" json:"avatar"`
	Country          int64     `gorm:"column:country" json:"country"`
	CreateAt         time.Time `gorm:"column:create_at" json:"create_at"`
	Email            string    `gorm:"column:email" json:"email"`
	EmailAuth        int64     `gorm:"column:email_auth" json:"email_auth"`
	GoogleAuthSwitch int64     `gorm:"column:google_auth_switch" json:"google_auth_switch"`
	ID               int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	IdentityAuth     int64     `gorm:"column:identity_auth" json:"identity_auth"`
	Language         string    `gorm:"column:language" json:"language"`
	Mobile           string    `gorm:"column:mobile" json:"mobile"`
	Password         string    `gorm:"column:password" json:"password"`
	RecommendCode    int64     `gorm:"column:recommend_code" json:"recommend_code"`
	Sex              int64     `gorm:"column:sex" json:"sex"`
	Status           int64     `gorm:"column:status" json:"status"`
	SwitchOrder      int64     `gorm:"column:switch_order" json:"switch_order"`
	Token            string    `gorm:"column:token" json:"token"`
	Type             int64     `gorm:"column:type" json:"type"`
	UpdateAt         time.Time `gorm:"column:update_at" json:"update_at"`
	Username         string    `gorm:"column:username" json:"username"`
	UUID             int64     `gorm:"column:uuid" json:"uuid"`
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "member"
}

func (e *Member) GetList() (Menu []*Member, err error) {

	table := orm.Eloquent.Table(e.TableName())
	// table = table.Where("menu_id = ?", e.MenuId)
	if err = table.Find(&Menu).Error; err != nil {
		return
	}
	return
}
