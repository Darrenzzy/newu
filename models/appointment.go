package models

import (
	orm "go-admin/global"
	"time"
)

type Appointment struct {
	City     string    `gorm:"column:city" json:"city"`
	Class    int64     `gorm:"column:class" json:"class"`
	Status   int64     `gorm:"column:status" json:"status"`
	Email    string    `gorm:"column:email" json:"email"`
	ID       int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Mobile   string    `gorm:"column:mobile" json:"mobile"`
	Name     string    `gorm:"column:name" json:"name"`
	Sex      string    `gorm:"column:sex" json:"sex"`
	UpdateBy time.Time `gorm:"column:update_by" json:"update_by"`
}

// TableName sets the insert table name for this struct type
func (a *Appointment) TableName() string {
	return "appointment"
}

func (e *Appointment) Get() (Appointment, error) {
	var doc Appointment
	table := orm.Eloquent.Table(e.TableName())
	if e.ID != 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err := table.Last(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
func (role *Appointment) GetPage(pageSize int, pageIndex int) ([]Appointment, int, error) {
	var doc []Appointment

	table := orm.Eloquent.Select("*").Model(role)

	var count int
	if err := table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}
func (e *Appointment) GetList() ([]*Appointment, error) {
	var doc []*Appointment
	table := orm.Eloquent.Table(e.TableName())
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *Appointment) Update() (update Appointment, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, e.ID).Error; err != nil {
		return
	}

	// 参数1:是要修改的数据
	// e.UpdateBy = time.Now().Format("2006-01-02 15:04:05")
	// 参数2:是修改的数据
	if err = orm.Eloquent.Model(&update).Save(&e).Error; err != nil {
		return
	}
	return
}

func (e *Appointment) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Appointment{}).Error; err != nil {
		return
	}
	Result = true
	return
}

// 添加
func (e Appointment) Insert() (id int64, err error) {
	e.UpdateBy = time.Now()
	e.Status = 1 // 待处理
	// 添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}
