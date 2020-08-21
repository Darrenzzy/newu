package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	orm "go-admin/global"
	"time"
)

type Member struct {
	// Avatar           string    `gorm:"column:avatar" json:"avatar"`
	// Country          int64     `gorm:"column:country" json:"country"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	Email    string    `gorm:"column:email" json:"email"`
	// EmailAuth        int64     `gorm:"column:email_auth" json:"email_auth"`
	// GoogleAuthSwitch int64     `gorm:"column:google_auth_switch" json:"google_auth_switch"`
	ID int64 `gorm:"column:id;primary_key" json:"id;primary_key"`
	// IdentityAuth     int64     `gorm:"column:identity_auth" json:"identity_auth"`
	// Language         string    `gorm:"column:language" json:"language"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Password string `gorm:"column:password" json:"password"`
	// RecommendCode    int64     `gorm:"column:recommend_code" json:"recommend_code"`
	// Sex              int64     `gorm:"column:sex" json:"sex"`
	// Status           int64     `gorm:"column:status" json:"status"`
	// SwitchOrder      int64     `gorm:"column:switch_order" json:"switch_order"`
	// Token            string    `gorm:"column:token" json:"token"`
	// Type             int64     `gorm:"column:type" json:"type"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
	Username string    `gorm:"column:username" json:"username"`
	Code     string    `gorm:"-" json:"code"`
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "member"
}

func (role *Member) GetPage(pageSize int, pageIndex int) ([]Member, int, error) {
	var doc []Member

	table := orm.Eloquent.Select("*").Table("member")
	if role.Username != "" {
		table = table.Where("username=?", role.Username)
	}

	if role.Mobile != "" {
		table = table.Where("mobile=?", role.Mobile)
	}

	var count int
	if err := table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	// table.Where("`deleted_at` IS NULL").Count(&count)
	table.Count(&count)
	return doc, count, nil
}

func (e *Member) GetList() (Menu []*Member, err error) {

	table := orm.Eloquent.Table(e.TableName())
	// table = table.Where("menu_id = ?", e.MenuId)
	if err = table.Find(&Menu).Error; err != nil {
		return
	}
	return
}

func (e *Member) Get() (data Member, err error) {
	table := orm.Eloquent.Table(e.TableName())
	if e.ID != 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err := table.Take(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (e *Member) Update() (update Member, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, e.ID).Error; err != nil {
		return
	}

	// 参数1:是要修改的数据
	// 参数2:是修改的数据
	if err = orm.Eloquent.Model(&update).Save(&e).Error; err != nil {
		return
	}
	return
}

func (e *Member) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Member{}).Error; err != nil {
		return
	}
	Result = true
	return
}

// 添加
func (e *Member) Insert() (id int64, err error) {
	var count int
	orm.Eloquent.Table(e.TableName()).Where("mobile = ?", e.Mobile).Count(&count)
	if count > 0 {
		err = errors.New("手机号已被注册！")
		return
	}
	e.Password = fmt.Sprintf("%x", md5.Sum([]byte(e.Password)))
	e.UpdateAt = time.Now()
	e.CreateAt = time.Now()
	// 添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

func (e *Member) Login() (err error) {
	mobile := e.Mobile
	pass := e.Password
	orm.Eloquent.Table(e.TableName()).Last(&e, "mobile=?", mobile)
	if e.ID == 0 {
		err = errors.New("用户不存在")
		return
	}
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(pass)))
	if md5Password != e.Password {
		err = errors.New("密码错误")
		return
	}
	e.UpdateAt = time.Now()
	if err = orm.Eloquent.Table(e.TableName()).Save(&e).Error; err != nil {
		return
	}

	return
}
