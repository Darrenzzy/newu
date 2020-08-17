package models

import orm "go-admin/global"

// 净值
type NetWorth struct {
	BuildBefore string `gorm:"column:build_before" json:"build_before"`
	Code        int    `gorm:"column:code" json:"code"`
	ID          int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	LastYear    string `gorm:"column:last_year" json:"last_year"`
	NetWorth    string `gorm:"column:net_worth" json:"net_worth"`
	NowYear     string `gorm:"column:now_year" json:"now_year"`
	SixMouth    string `gorm:"column:six_mouth" json:"six_mouth"`
	ThreeMuoth  string `gorm:"column:three_muoth" json:"three_muoth"`
	UnitWorth   string `gorm:"column:unit_worth" json:"unit_worth"`
	WondName    string `gorm:"column:wond_name" json:"wond_name"`
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

func (e *NetWorth) Update() (update NetWorth, err error) {
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
