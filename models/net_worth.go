package models

import (
	orm "go-admin/global"
	"time"
)

// 净值
type NetWorth struct {
	BuildBefore string `gorm:"column:build_before" json:"build_before"` // 成立以来(%)
	Code        int    `gorm:"column:code" json:"code"`                 // 基金代码
	ID          int64  `gorm:"column:id;primary_key" json:"id;primary_key"`
	LastYear    string `gorm:"column:last_year" json:"last_year"` // 近一年(%)
	NetWorth    string `gorm:"column:net_worth" json:"net_worth"`
	NowYear     string `gorm:"column:now_year" json:"now_year"`       // 今年以来(%)
	SixMouth    string `gorm:"column:six_mouth" json:"six_mouth"`     // 近六个月(%)
	ThreeMuoth  string `gorm:"column:three_muoth" json:"three_muoth"` // 近三个月(%)
	UnitWorth   string `gorm:"column:unit_worth" json:"unit_worth"`   // 单位净值
	WondName    string `gorm:"column:wond_name" json:"wond_name"`     // 基金名称
	DateWorth   string `gorm:"column:date_worth" json:"date_worth"`   // 净值日期
	Intro       string `gorm:"column:intro" json:"intro"`             // 简介
	Intro2      string `gorm:"column:intro2" json:"intto2"`           // 简介 2
	IsLimit     int    `gorm:"column:is_limit" json:"is_limit"`       // 是否限制
	CreateBy    string `json:"create_by" gorm:"size:128;"`            //
	UpdateBy    string `json:"update_by" gorm:"size:128;"`
}

// TableName sets the insert table name for this struct type
func (n *NetWorth) TableName() string {
	return "net_worth"
}

func (e *NetWorth) Get() (NetWorth, error) {
	var doc NetWorth
	table := orm.Eloquent.Table(e.TableName())
	if e.ID != 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err := table.Last(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
func (role *NetWorth) GetPage(pageSize int, pageIndex int) ([]NetWorth, int, error) {
	var doc []NetWorth

	table := orm.Eloquent.Select("*").Table("net_worth")

	var count int
	if err := table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}
func (e *NetWorth) GetList() ([]*NetWorth, error) {
	var doc []*NetWorth
	table := orm.Eloquent.Table(e.TableName())
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *NetWorth) Update() (update NetWorth, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, e.ID).Error; err != nil {
		return
	}

	// 参数1:是要修改的数据
	e.UpdateBy = time.Now().Format("2006-01-02 15:04:05")
	// 参数2:是修改的数据
	if err = orm.Eloquent.Model(&update).Save(&e).Error; err != nil {
		return
	}
	return
}

func (e *NetWorth) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&NetWorth{}).Error; err != nil {
		return
	}
	Result = true
	return
}

// 添加
func (e NetWorth) Insert() (id int64, err error) {

	e.CreateBy = time.Now().Format("2006-01-02 15:04:05")
	// 添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}
