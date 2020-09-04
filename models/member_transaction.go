package models

import (
	"fmt"
	orm "go-admin/global"
	"strconv"
	"time"
)

type MemberTransaction struct {
	Amount      float64   `gorm:"column:amount" json:"amount"`
	CoinID      int       `gorm:"column:coin_id" json:"coin_id"`
	CreateAt    time.Time `gorm:"column:create_time" json:"create_time"`
	Fee         float64   `gorm:"column:fee" json:"fee"`
	ID          int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID    int       `gorm:"column:member_id" json:"member_id"`
	OtherAmount float64   `gorm:"column:other_amount" json:"other_amount"`
	Rate        float64   `gorm:"column:rate" json:"rate"`
	Remark      string    `gorm:"column:remark" json:"remark"`
	Total       float64   `gorm:"column:total" json:"total"`
	Type        int       `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (m *MemberTransaction) TableName() string {
	return "member_transaction"
}

func (e *MemberTransaction) GetPage(pageSize int, pageIndex int) ([]MemberTransaction, int, error) {
	var doc []MemberTransaction

	table := orm.Eloquent.Select("*").Model(e)
	// if role.Username != "" {
	// 	table = table.Where("username=?", role.Username)
	// }
	//
	// if role.Mobile != "" {
	// 	table = table.Where("mobile=?", role.Mobile)
	// }

	var count int
	if err := table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	// table.Where("`deleted_at` IS NULL").Count(&count)
	table.Count(&count)
	return doc, count, nil
}

func (e *MemberTransaction) GetList() (Menu []*MemberTransaction, err error) {

	table := orm.Eloquent.Table(e.TableName())
	// table = table.Where("menu_id = ?", e.MenuId)
	if err = table.Find(&Menu).Error; err != nil {
		return
	}
	return
}

func (e *MemberTransaction) Get() (data MemberTransaction, err error) {
	table := orm.Eloquent.Table(e.TableName())
	if e.ID != 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err := table.Take(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (e *MemberTransaction) Update() (update MemberTransaction, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, e.ID).Error; err != nil {
		return
	}
	e.CheckAmount()

	// 参数1:是要修改的数据
	// 参数2:是修改的数据
	if err = orm.Eloquent.Model(&update).Save(&e).Error; err != nil {
		return
	}
	return
}

func (e *MemberTransaction) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&MemberTransaction{}).Error; err != nil {
		return
	}
	Result = true
	return
}

// 添加
func (e *MemberTransaction) Insert() (id int64, err error) {
	e.CheckAmount()
	if e.CreateAt.Unix() < 0 {
		e.CreateAt = time.Now()
	}
	// 添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

func (e *MemberTransaction) CheckAmount() {
	var last MemberTransaction
	DD := orm.Eloquent.Table(e.TableName())
	if e.ID > 0 {
		DD = DD.Where("id<?", e.ID)
	}
	DD.Last(&last)
	// 	rate cny
	if e.OtherAmount != 0 {
		e.Amount, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", e.OtherAmount/e.Rate), 64)
	} else if e.Amount != 0 {
		// usdt * rate
		e.OtherAmount = e.Amount * e.Rate
	}
	// 	total
	if e.Type == 1 {
		e.Total = last.Total + e.Amount
	} else {
		e.Total = last.Total - e.Amount

	}

}
